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
import Link from 'next/link';

interface Subject {
    id: string;
    name: string;
    code: string;
}

interface Venue {
    id: string;
    name: string;
}

interface Exam {
    id: string;
    title: string;
    examType: string;
    examDate: string; // YYYY-MM-DD
    startTime?: string; // HH:MM:SS
    endTime?: string; // HH:MM:SS
    durationMinutes?: number;
    prepStatus: string;
    subjectId?: string;
    venueId?: string;
    syllabusUnits?: string[];
    syllabusTopics?: string[];
    syllabusNotes?: string;
    // ... other fields as needed for display
}

export default function ExamsPage() {
    const { token } = useAuthStore();
    const [exams, setExams] = useState<Exam[]>([]);
    const [subjects, setSubjects] = useState<Subject[]>([]);
    const [venues, setVenues] = useState<Venue[]>([]);
    const [newExam, setNewExam] = useState({
        title: '',
        examType: 'cat1',
        examDate: '', // YYYY-MM-DD
        startTime: '',
        endTime: '',
        durationMinutes: 60,
        prepStatus: 'not_started',
        subjectId: '',
        venueId: '',
        syllabusUnits: [],
        syllabusTopics: [],
        syllabusNotes: '',
    });
    const [isLoading, setIsLoading] = useState(false);

    useEffect(() => {
        if (token) {
            fetchExams();
            fetchDependencies();
        }
    }, [token]);

    const fetchExams = async () => {
        try {
            const res = await fetch("/api/exams", {
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            });
            if (res.ok) {
                const data = await res.json();
                setExams(data);
            } else {
                toast.error("Failed to fetch exams.");
            }
        } catch (error) {
            toast.error("An unexpected error occurred while fetching exams.");
            console.error("Fetch exams error:", error);
        }
    };

    const fetchDependencies = async () => {
        try {
            const [subjectsRes, venuesRes] = await Promise.all([
                fetch("/api/subjects", { headers: { Authorization: `Bearer ${token}` } }),
                fetch("/api/venues", { headers: { Authorization: `Bearer ${token}` } }),
            ]);

            if (subjectsRes.ok) setSubjects(await subjectsRes.json());
            else toast.error("Failed to fetch subjects.");
            if (venuesRes.ok) setVenues(await venuesRes.json());
            else toast.error("Failed to fetch venues.");

        } catch (error) {
            toast.error("Error fetching exam dependencies.");
            console.error("Fetch dependencies error:", error);
        }
    };

    const handleInputChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        const { name, value, type } = e.target;
        setNewExam(prev => ({
            ...prev,
            [name]: type === 'number' ? parseInt(value) : value,
        }));
    };

    const handleSelectChange = (name: string, value: string) => {
        setNewExam(prev => ({ ...prev, [name]: value }));
    };

    const handleSubmitNewExam = async (e: React.FormEvent) => {
        e.preventDefault();
        setIsLoading(true);

        const payload = {
            ...newExam,
            subjectId: newExam.subjectId || null,
            venueId: newExam.venueId || null,
            startTime: newExam.startTime || null,
            endTime: newExam.endTime || null,
            durationMinutes: newExam.durationMinutes || null,
            syllabusUnits: newExam.syllabusUnits.filter(Boolean), // Remove empty strings
            syllabusTopics: newExam.syllabusTopics.filter(Boolean),
            syllabusNotes: newExam.syllabusNotes || null,
        };

        try {
            const res = await fetch("/api/exams", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: `Bearer ${token}`,
                },
                body: JSON.stringify(payload),
            });

            if (res.ok) {
                toast.success("Exam created successfully!");
                setNewExam({ // Reset form
                    title: '',
                    examType: 'cat1',
                    examDate: '',
                    startTime: '',
                    endTime: '',
                    durationMinutes: 60,
                    prepStatus: 'not_started',
                    subjectId: '',
                    venueId: '',
                    syllabusUnits: [],
                    syllabusTopics: [],
                    syllabusNotes: '',
                });
                fetchExams(); // Refresh list
            } else {
                const errorData = await res.json();
                toast.error(errorData.error || "Failed to create exam.");
            }
        } catch (error) {
            toast.error("An unexpected error occurred while creating exam.");
            console.error("Create exam error:", error);
        } finally {
            setIsLoading(false);
        }
    };

    const getPrepStatusColor = (status: string) => {
        switch (status) {
            case 'not_started': return 'text-gray-500';
            case 'in_progress': return 'text-blue-500';
            case 'revision': return 'text-yellow-500';
            case 'ready': return 'text-green-500';
            default: return 'text-gray-500';
        }
    };

    return (
        <div className="space-y-6">
            <Toaster position="top-center" richColors />

            <Card className="bg-gray-800 border-gray-700">
                <CardHeader>
                    <CardTitle className="text-white">Create New Exam</CardTitle>
                </CardHeader>
                <CardContent>
                    <form onSubmit={handleSubmitNewExam} className="grid grid-cols-1 md:grid-cols-2 gap-4">
                        <Input name="title" placeholder="Exam Title (e.g., CAT-1, FAT)" value={newExam.title} onChange={handleInputChange} required className="bg-gray-700 border-gray-600 text-white col-span-2" />
                        
                        <div>
                            <label className="block text-gray-300 text-sm font-bold mb-2">Exam Type</label>
                            <Select name="examType" value={newExam.examType} onValueChange={(val) => handleSelectChange('examType', val)} required>
                                <SelectTrigger className="w-full bg-gray-700 border-gray-600 text-white">
                                    <SelectValue placeholder="Select type" />
                                </SelectTrigger>
                                <SelectContent className="bg-gray-700 border-gray-600 text-white">
                                    <SelectItem value="cat1">CAT-1</SelectItem>
                                    <SelectItem value="cat2">CAT-2</SelectItem>
                                    <SelectItem value="cat3">CAT-3</SelectItem>
                                    <SelectItem value="fat">FAT</SelectItem>
                                    <SelectItem value="model">Model Exam</SelectItem>
                                    <SelectItem value="quiz">Quiz</SelectItem>
                                    <SelectItem value="viva">Viva</SelectItem>
                                    <SelectItem value="practical">Practical</SelectItem>
                                    <SelectItem value="other">Other</SelectItem>
                                </SelectContent>
                            </Select>
                        </div>

                        <Input name="examDate" type="date" placeholder="Exam Date" value={newExam.examDate} onChange={handleInputChange} required className="bg-gray-700 border-gray-600 text-white" />
                        <Input name="startTime" type="time" placeholder="Start Time (Optional)" value={newExam.startTime} onChange={handleInputChange} className="bg-gray-700 border-gray-600 text-white" />
                        <Input name="endTime" type="time" placeholder="End Time (Optional)" value={newExam.endTime} onChange={handleInputChange} className="bg-gray-700 border-gray-600 text-white" />
                        <Input name="durationMinutes" type="number" placeholder="Duration (minutes, Optional)" value={newExam.durationMinutes} onChange={handleInputChange} className="bg-gray-700 border-gray-600 text-white" />
                        
                        <div>
                            <label className="block text-gray-300 text-sm font-bold mb-2">Subject (Optional)</label>
                            <Select name="subjectId" value={newExam.subjectId || ''} onValueChange={(val) => handleSelectChange('subjectId', val)}>
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
                            <label className="block text-gray-300 text-sm font-bold mb-2">Venue (Optional)</label>
                            <Select name="venueId" value={newExam.venueId || ''} onValueChange={(val) => handleSelectChange('venueId', val)}>
                                <SelectTrigger className="w-full bg-gray-700 border-gray-600 text-white">
                                    <SelectValue placeholder="Select venue" />
                                </SelectTrigger>
                                <SelectContent className="bg-gray-700 border-gray-600 text-white">
                                    {venues.map(venue => (
                                        <SelectItem key={venue.id} value={venue.id}>{venue.name}</SelectItem>
                                    ))}
                                </SelectContent>
                            </Select>
                        </div>

                        <Input name="syllabusUnits" placeholder="Syllabus Units (comma-separated)" value={newExam.syllabusUnits.join(', ')} onChange={(e) => setNewExam(prev => ({ ...prev, syllabusUnits: e.target.value.split(',').map(s => s.trim()) }))} className="bg-gray-700 border-gray-600 text-white col-span-2" />
                        <Input name="syllabusTopics" placeholder="Syllabus Topics (comma-separated)" value={newExam.syllabusTopics.join(', ')} onChange={(e) => setNewExam(prev => ({ ...prev, syllabusTopics: e.target.value.split(',').map(s => s.trim()) }))} className="bg-gray-700 border-gray-600 text-white col-span-2" />
                        <Textarea name="syllabusNotes" placeholder="Syllabus Notes (Optional)" value={newExam.syllabusNotes} onChange={handleInputChange} className="bg-gray-700 border-gray-600 text-white col-span-2" />
                        
                        <div>
                            <label className="block text-gray-300 text-sm font-bold mb-2">Preparation Status</label>
                            <Select name="prepStatus" value={newExam.prepStatus} onValueChange={(val) => handleSelectChange('prepStatus', val)} required>
                                <SelectTrigger className="w-full bg-gray-700 border-gray-600 text-white">
                                    <SelectValue placeholder="Select status" />
                                </SelectTrigger>
                                <SelectContent className="bg-gray-700 border-gray-600 text-white">
                                    <SelectItem value="not_started">Not Started</SelectItem>
                                    <SelectItem value="in_progress">In Progress</SelectItem>
                                    <SelectItem value="revision">Revision</SelectItem>
                                    <SelectItem value="ready">Ready</SelectItem>
                                </SelectContent>
                            </Select>
                        </div>

                        <div className="md:col-span-2">
                            <Button type="submit" disabled={isLoading} className="w-full bg-blue-600 hover:bg-blue-700">
                                {isLoading ? "Creating..." : "Add Exam"}
                            </Button>
                        </div>
                    </form>
                </CardContent>
            </Card>

            <Card className="bg-gray-800 border-gray-700">
                <CardHeader>
                    <CardTitle className="text-white">Your Exams</CardTitle>
                </CardHeader>
                <CardContent>
                    {exams.length === 0 ? (
                        <p className="text-gray-400">No exams found. Create one above!</p>
                    ) : (
                        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                            {exams.map(exam => (
                                <Card key={exam.id} className="bg-gray-700 border-gray-600 p-4">
                                    <h3 className="font-bold text-lg text-white">{exam.title} ({exam.examType})</h3>
                                    <p className="text-gray-300 text-sm">Date: {new Date(exam.examDate).toLocaleDateString()}</p>
                                    {exam.startTime && <p className="text-gray-300 text-sm">Time: {exam.startTime} - {exam.endTime}</p>}
                                    {exam.subjectId && <p className="text-gray-300 text-sm">Subject: {subjects.find(sub => sub.id === exam.subjectId)?.name}</p>}
                                    {exam.venueId && <p className="text-gray-300 text-sm">Venue: {venues.find(v => v.id === exam.venueId)?.name}</p>}
                                    <p className={cn("text-sm font-semibold mt-2", getPrepStatusColor(exam.prepStatus))}>Prep Status: {exam.prepStatus}</p>
                                </Card>
                            ))}
                        </div>
                    )}
                </CardContent>
            </Card>
        </div>
    );
}
