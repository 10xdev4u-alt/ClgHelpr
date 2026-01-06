"use client";

import { useEffect, useState } from 'react';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { Textarea } from '@/components/ui/textarea';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { Toaster, toast } from 'sonner';
import { useAuthStore } from '@/stores/auth-store';
import { cn } from '@/lib/utils';
import { useParams } from 'next/navigation';

interface StudySession {
    id: string;
    studyPlanId: string;
    subjectId?: string;
    plannedStartTime?: string; // ISO string
    plannedEndTime?: string;   // ISO string
    sessionType: string;
    topicsToCover?: string[];
    topicsCovered?: string[];
    status: string;
    completionPercentage: number;
    // ... other fields as needed
}

interface Subject {
    id: string;
    name: string;
    code: string;
}

export default function StudySessionsPage() {
    const { planId } = useParams();
    const { token } = useAuthStore();
    const [studySessions, setStudySessions] = useState<StudySession[]>([]);
    const [subjects, setSubjects] = useState<Subject[]>([]);
    const [newSession, setNewSession] = useState({
        studyPlanId: planId as string,
        subjectId: '',
        plannedStartTime: '', // YYYY-MM-DDTHH:MM
        plannedEndTime: '',   // YYYY-MM-DDTHH:MM
        sessionType: 'study',
        topicsToCover: '', // comma-separated
        topicsCovered: '', // comma-separated
        status: 'planned',
        completionPercentage: 0,
    });
    const [isLoading, setIsLoading] = useState(false);

    useEffect(() => {
        if (token && planId) {
            fetchStudySessions();
            fetchSubjects();
        }
    }, [token, planId]);

    const fetchStudySessions = async () => {
        try {
            const res = await fetch(`/api/study-sessions/plan/${planId}`, {
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            });
            if (res.ok) {
                const data = await res.json();
                setStudySessions(data);
            } else {
                toast.error("Failed to fetch study sessions.");
            }
        } catch (error) {
            toast.error("An unexpected error occurred while fetching study sessions.");
            console.error("Fetch study sessions error:", error);
        }
    };

    const fetchSubjects = async () => {
        try {
            const subjectsRes = await fetch("/api/subjects", { headers: { Authorization: `Bearer ${token}` } });
            if (subjectsRes.ok) setSubjects(await subjectsRes.json());
            else toast.error("Failed to fetch subjects.");

        } catch (error) {
            toast.error("Error fetching subject dependencies.");
            console.error("Fetch dependencies error:", error);
        }
    };

    const handleInputChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        const { name, value, type } = e.target;
        setNewSession(prev => ({
            ...prev,
            [name]: type === 'number' ? parseInt(value) : value,
        }));
    };

    const handleSelectChange = (name: string, value: string) => {
        setNewSession(prev => ({ ...prev, [name]: value }));
    };

    const handleSubmitNewSession = async (e: React.FormEvent) => {
        e.preventDefault();
        setIsLoading(true);

        const payload = {
            ...newSession,
            studyPlanId: planId as string,
            subjectId: newSession.subjectId || null,
            plannedStartTime: newSession.plannedStartTime ? new Date(newSession.plannedStartTime).toISOString() : null,
            plannedEndTime: newSession.plannedEndTime ? new Date(newSession.plannedEndTime).toISOString() : null,
            topicsToCover: newSession.topicsToCover ? newSession.topicsToCover.split(',').map(s => s.trim()).filter(Boolean) : [],
            topicsCovered: newSession.topicsCovered ? newSession.topicsCovered.split(',').map(s => s.trim()).filter(Boolean) : [],
        };

        try {
            const res = await fetch("/api/study-sessions", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: `Bearer ${token}`,
                },
                body: JSON.stringify(payload),
            });

            if (res.ok) {
                toast.success("Study Session created successfully!");
                setNewSession({ // Reset form
                    studyPlanId: planId as string,
                    subjectId: '',
                    plannedStartTime: '',
                    plannedEndTime: '',
                    sessionType: 'study',
                    topicsToCover: '',
                    topicsCovered: '',
                    status: 'planned',
                    completionPercentage: 0,
                });
                fetchStudySessions(); // Refresh list
            } else {
                const errorData = await res.json();
                toast.error(errorData.error || "Failed to create study session.");
            }
        } catch (error) {
            toast.error("An unexpected error occurred while creating session.");
            console.error("Create session error:", error);
        } finally {
            setIsLoading(false);
        }
    };

    const getStatusColor = (status: string) => {
        switch (status) {
            case 'planned': return 'text-blue-400';
            case 'in_progress': return 'text-yellow-500';
            case 'completed': return 'text-green-500';
            case 'partial': return 'text-orange-400';
            case 'skipped': return 'text-red-500';
            default: return 'text-gray-500';
        }
    };

    return (
        <div className="space-y-6">
            <Toaster position="top-center" richColors />

            <Card className="bg-gray-800 border-gray-700">
                <CardHeader>
                    <CardTitle className="text-white">Create New Study Session for Plan: {planId}</CardTitle>
                </CardHeader>
                <CardContent>
                    <form onSubmit={handleSubmitNewSession} className="grid grid-cols-1 md:grid-cols-2 gap-4">
                        <div>
                            <label className="block text-gray-300 text-sm font-bold mb-2">Subject (Optional)</label>
                            <Select name="subjectId" value={newSession.subjectId || ''} onValueChange={(val) => handleSelectChange('subjectId', val)}>
                                <SelectTrigger className="w-full bg-gray-700 border-gray-600 text-white">
                                    <SelectValue placeholder="Select subject" />
                                </SelectTrigger>
                                <SelectContent className="bg-gray-700 border-gray-600 text-white">
                                    {subjects.map(sub => (
                                        <SelectItem key={sub.id} value={sub.id}>{sub.name} ({sub.code})</SelectItem>
                                    ))}
                                </SelectContent>
                            </Select>
                        </div>
                        
                        <div>
                            <label className="block text-gray-300 text-sm font-bold mb-2">Session Type</label>
                            <Select name="sessionType" value={newSession.sessionType} onValueChange={(val) => handleSelectChange('sessionType', val)} required>
                                <SelectTrigger className="w-full bg-gray-700 border-gray-600 text-white">
                                    <SelectValue placeholder="Select type" />
                                </SelectTrigger>
                                <SelectContent className="bg-gray-700 border-gray-600 text-white">
                                    <SelectItem value="study">Study</SelectItem>
                                    <SelectItem value="revision">Revision</SelectItem>
                                    <SelectItem value="practice">Practice</SelectItem>
                                    <SelectItem value="assignment">Assignment</SelectItem>
                                    <SelectItem value="lab_prep">Lab Prep</SelectItem>
                                    <SelectItem value="exam_prep">Exam Prep</SelectItem>
                                </SelectContent>
                            </Select>
                        </div>

                        <Input name="plannedStartTime" type="datetime-local" placeholder="Planned Start Time (Optional)" value={newSession.plannedStartTime} onChange={handleInputChange} className="bg-gray-700 border-gray-600 text-white" />
                        <Input name="plannedEndTime" type="datetime-local" placeholder="Planned End Time (Optional)" value={newSession.plannedEndTime} onChange={handleInputChange} className="bg-gray-700 border-gray-600 text-white" />
                        
                        <Input name="topicsToCover" placeholder="Topics to Cover (comma-separated)" value={newSession.topicsToCover} onChange={handleInputChange} className="bg-gray-700 border-gray-600 text-white col-span-2" />
                        <Input name="topicsCovered" placeholder="Topics Covered (comma-separated)" value={newSession.topicsCovered} onChange={handleInputChange} className="bg-gray-700 border-gray-600 text-white col-span-2" />

                        <div>
                            <label className="block text-gray-300 text-sm font-bold mb-2">Status</label>
                            <Select name="status" value={newSession.status} onValueChange={(val) => handleSelectChange('status', val)} required>
                                <SelectTrigger className="w-full bg-gray-700 border-gray-600 text-white">
                                    <SelectValue placeholder="Select status" />
                                </SelectTrigger>
                                <SelectContent className="bg-gray-700 border-gray-600 text-white">
                                    <SelectItem value="planned">Planned</SelectItem>
                                    <SelectItem value="in_progress">In Progress</SelectItem>
                                    <SelectItem value="completed">Completed</SelectItem>
                                    <SelectItem value="partial">Partial</SelectItem>
                                    <SelectItem value="skipped">SkiSelectItem</SelectItem>
                                </SelectContent>
                            </Select>
                        </div>
                        <Input name="completionPercentage" type="number" placeholder="Completion % (0-100)" value={newSession.completionPercentage || ''} onChange={handleInputChange} className="bg-gray-700 border-gray-600 text-white" />
                        
                        <div className="md:col-span-2">
                            <Button type="submit" disabled={isLoading} className="w-full bg-blue-600 hover:bg-blue-700">
                                {isLoading ? "Creating..." : "Add Study Session"}
                            </Button>
                        </div>
                    </form>
                </CardContent>
            </Card>

            <Card className="bg-gray-800 border-gray-700">
                <CardHeader>
                    <CardTitle className="text-white">Study Sessions for Plan: {planId}</CardTitle>
                </CardHeader>
                <CardContent>
                    {studySessions.length === 0 ? (
                        <p className="text-gray-400">No study sessions found for this plan. Create one above!</p>
                    ) : (
                        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                            {studySessions.map(session => (
                                <Card key={session.id} className="bg-gray-700 border-gray-600 p-4">
                                    <h3 className="font-bold text-lg text-white">{session.sessionType}</h3>
                                    {session.subjectId && <p className="text-gray-300 text-sm">Subject: {subjects.find(sub => sub.id === session.subjectId)?.name}</p>}
                                    {session.plannedStartTime && <p className="text-gray-300 text-sm">Start: {new Date(session.plannedStartTime).toLocaleString()}</p>}
                                    {session.plannedEndTime && <p className="text-gray-300 text-sm">End: {new Date(session.plannedEndTime).toLocaleString()}</p>}
                                    <p className={cn("text-sm font-semibold mt-2", getStatusColor(session.status))}>Status: {session.status} ({session.completionPercentage}%)</p>
                                </Card>
                            ))}
                        </div>
                    )}
                </CardContent>
            </Card>
        </div>
    );
}
