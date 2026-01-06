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

interface StudyPlan {
    id: string;
    title: string;
    planDate: string; // YYYY-MM-DD
    planType: string;
    status: string;
    notes?: string;
}

interface Subject {
    id: string;
    name: string;
    code: string;
}

export default function StudyPlansPage() {
    const { token } = useAuthStore();
    const [studyPlans, setStudyPlans] = useState<StudyPlan[]>([]);
    const [newStudyPlan, setNewStudyPlan] = useState({
        title: '',
        planDate: '', // YYYY-MM-DD
        planType: 'daily',
        notes: '',
        status: 'planned',
    });
    const [isLoading, setIsLoading] = useState(false);

    useEffect(() => {
        if (token) {
            fetchStudyPlans();
        }
    }, [token]);

    const fetchStudyPlans = async () => {
        try {
            const res = await fetch("/api/study-plans", {
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            });
            if (res.ok) {
                const data = await res.json();
                setStudyPlans(data);
            } else {
                toast.error("Failed to fetch study plans.");
            }
        } catch (error) {
            toast.error("An unexpected error occurred while fetching study plans.");
            console.error("Fetch study plans error:", error);
        }
    };

    const handleInputChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        const { name, value, type } = e.target;
        setNewStudyPlan(prev => ({
            ...prev,
            [name]: type === 'number' ? parseInt(value) : value,
        }));
    };

    const handleSelectChange = (name: string, value: string) => {
        setNewStudyPlan(prev => ({ ...prev, [name]: value }));
    };

    const handleSubmitNewStudyPlan = async (e: React.FormEvent) => {
        e.preventDefault();
        setIsLoading(true);

        const payload = {
            ...newStudyPlan,
            notes: newStudyPlan.notes || null,
        };

        try {
            const res = await fetch("/api/study-plans", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: `Bearer ${token}`,
                },
                body: JSON.stringify(payload),
            });

            if (res.ok) {
                toast.success("Study Plan created successfully!");
                setNewStudyPlan({ // Reset form
                    title: '',
                    planDate: '',
                    planType: 'daily',
                    notes: '',
                    status: 'planned',
                });
                fetchStudyPlans(); // Refresh list
            } else {
                const errorData = await res.json();
                toast.error(errorData.error || "Failed to create study plan.");
            }
        } catch (error) {
            toast.error("An unexpected error occurred while creating study plan.");
            console.error("Create study plan error:", error);
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
                    <CardTitle className="text-white">Create New Study Plan</CardTitle>
                </CardHeader>
                <CardContent>
                    <form onSubmit={handleSubmitNewStudyPlan} className="grid grid-cols-1 md:grid-cols-2 gap-4">
                        <Input name="title" placeholder="Plan Title (e.g., Saturday Revision)" value={newStudyPlan.title} onChange={handleInputChange} required className="bg-gray-700 border-gray-600 text-white col-span-2" />
                        
                        <Input name="planDate" type="date" placeholder="Plan Date" value={newStudyPlan.planDate} onChange={handleInputChange} required className="bg-gray-700 border-gray-600 text-white" />
                        
                        <div>
                            <label className="block text-gray-300 text-sm font-bold mb-2">Plan Type</label>
                            <Select name="planType" value={newStudyPlan.planType} onValueChange={(val) => handleSelectChange('planType', val)} required>
                                <SelectTrigger className="w-full bg-gray-700 border-gray-600 text-white">
                                    <SelectValue placeholder="Select type" />
                                </SelectTrigger>
                                <SelectContent className="bg-gray-700 border-gray-600 text-white">
                                    <SelectItem value="daily">Daily</SelectItem>
                                    <SelectItem value="weekly">Weekly</SelectItem>
                                    <SelectItem value="weekend">Weekend</SelectItem>
                                    <SelectItem value="exam_prep">Exam Prep</SelectItem>
                                    <SelectItem value="revision">Revision</SelectItem>
                                    <SelectItem value="custom">Custom</SelectItem>
                                </SelectContent>
                            </Select>
                        </div>

                        <div>
                            <label className="block text-gray-300 text-sm font-bold mb-2">Status</label>
                            <Select name="status" value={newStudyPlan.status} onValueChange={(val) => handleSelectChange('status', val)} required>
                                <SelectTrigger className="w-full bg-gray-700 border-gray-600 text-white">
                                    <SelectValue placeholder="Select status" />
                                </SelectTrigger>
                                <SelectContent className="bg-gray-700 border-gray-600 text-white">
                                    <SelectItem value="planned">Planned</SelectItem>
                                    <SelectItem value="in_progress">In Progress</SelectItem>
                                    <SelectItem value="completed">Completed</SelectItem>
                                    <SelectItem value="partial">Partial</SelectItem>
                                    <SelectItem value="skipped">Skipped</SelectItem>
                                </SelectContent>
                            </Select>
                        </div>
                        <Textarea name="notes" placeholder="Notes (Optional)" value={newStudyPlan.notes} onChange={handleInputChange} className="bg-gray-700 border-gray-600 text-white col-span-2" />
                        
                        <div className="md:col-span-2">
                            <Button type="submit" disabled={isLoading} className="w-full bg-blue-600 hover:bg-blue-700">
                                {isLoading ? "Creating..." : "Add Study Plan"}
                            </Button>
                        </div>
                    </form>
                </CardContent>
            </Card>

            <Card className="bg-gray-800 border-gray-700">
                <CardHeader>
                    <CardTitle className="text-white">Your Study Plans</CardTitle>
                </CardHeader>
                <CardContent>
                    {studyPlans.length === 0 ? (
                        <p className="text-gray-400">No study plans found. Create one above!</p>
                    ) : (
                        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                            {studyPlans.map(plan => (
                                <Card key={plan.id} className="bg-gray-700 border-gray-600 p-4">
                                    <h3 className="font-bold text-lg text-white">{plan.title}</h3>
                                    <p className="text-gray-300 text-sm">Date: {new Date(plan.planDate).toLocaleDateString()}</p>
                                    <p className="text-gray-300 text-sm">Type: {plan.planType}</p>
                                    <p className={cn("text-sm font-semibold mt-2", getStatusColor(plan.status))}>Status: {plan.status}</p>
                                    {plan.notes && <p className="text-gray-300 text-sm mt-1">Notes: {plan.notes}</p>}
                                    <Link href={`/dashboard/study/${plan.id}`} className="text-blue-400 hover:underline text-sm mt-2 block">
                                        View Sessions
                                    </Link>
                                </Card>
                            ))}
                        </div>
                    )}
                </CardContent>
            </Card>
        </div>
    );
}
