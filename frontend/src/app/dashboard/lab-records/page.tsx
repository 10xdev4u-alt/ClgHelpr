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

interface LabRecord {
    id: string;
    experimentNumber: number;
    title: string;
    status: string;
    labDate?: string; // YYYY-MM-DD
    subjectId?: string;
    // ... other fields as needed for display
}

export default function LabRecordsPage() {
    const { token } = useAuthStore();
    const [labRecords, setLabRecords] = useState<LabRecord[]>([]);
    const [subjects, setSubjects] = useState<Subject[]>([]);
    const [newLabRecord, setNewLabRecord] = useState({
        experimentNumber: 0,
        title: '',
        status: 'pending',
        labDate: '', // YYYY-MM-DD
        subjectId: '',
        aim: '',
        algorithm: '',
        code: '',
        output: '',
        observations: '',
        result: '',
        vivaQuestions: '', // comma-separated
        printRequired: true,
        pagesToPrint: 0,
        marks: 0,
        staffRemarks: '',
    });
    const [isLoading, setIsLoading] = useState(false);

    useEffect(() => {
        if (token) {
            fetchLabRecords();
            fetchDependencies();
        }
    }, [token]);

    const fetchLabRecords = async () => {
        try {
            const res = await fetch("/api/lab-records", {
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            });
            if (res.ok) {
                const data = await res.json();
                setLabRecords(data);
            } else {
                toast.error("Failed to fetch lab records.");
            }
        } catch (error) {
            toast.error("An unexpected error occurred while fetching lab records.");
            console.error("Fetch lab records error:", error);
        }
    };

    const fetchDependencies = async () => {
        try {
            const subjectsRes = await fetch("/api/subjects", { headers: { Authorization: `Bearer ${token}` } });
            if (subjectsRes.ok) setSubjects(await subjectsRes.json());
            else toast.error("Failed to fetch subjects.");

        } catch (error) {
            toast.error("Error fetching lab record dependencies.");
            console.error("Fetch dependencies error:", error);
        }
    };

    const handleInputChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        const { name, value, type, checked } = e.target;
        setNewLabRecord(prev => ({
            ...prev,
            [name]: type === 'checkbox' ? checked : (type === 'number' ? parseInt(value) : value),
        }));
    };

    const handleSelectChange = (name: string, value: string) => {
        setNewLabRecord(prev => ({ ...prev, [name]: value }));
    };

    const handleSubmitNewLabRecord = async (e: React.FormEvent) => {
        e.preventDefault();
        setIsLoading(true);

        const payload = {
            ...newLabRecord,
            subjectId: newLabRecord.subjectId || null,
            labDate: newLabRecord.labDate || null,
            aim: newLabRecord.aim || null,
            algorithm: newLabRecord.algorithm || null,
            code: newLabRecord.code || null,
            output: newLabRecord.output || null,
            observations: newLabRecord.observations || null,
            result: newLabRecord.result || null,
            vivaQuestions: newLabRecord.vivaQuestions.split(',').map(s => s.trim()).filter(Boolean),
            pagesToPrint: newLabRecord.pagesToPrint || null,
            marks: newLabRecord.marks || null,
            staffRemarks: newLabRecord.staffRemarks || null,
        };

        try {
            const res = await fetch("/api/lab-records", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: `Bearer ${token}`,
                },
                body: JSON.stringify(payload),
            });

            if (res.ok) {
                toast.success("Lab Record created successfully!");
                setNewLabRecord({ // Reset form
                    experimentNumber: 0,
                    title: '',
                    status: 'pending',
                    labDate: '',
                    subjectId: '',
                    aim: '',
                    algorithm: '',
                    code: '',
                    output: '',
                    observations: '',
                    result: '',
                    vivaQuestions: '',
                    printRequired: true,
                    pagesToPrint: 0,
                    marks: 0,
                    staffRemarks: '',
                });
                fetchLabRecords(); // Refresh list
            } else {
                const errorData = await res.json();
                toast.error(errorData.error || "Failed to create lab record.");
            }
        } catch (error) {
            toast.error("An unexpected error occurred while creating lab record.");
            console.error("Create lab record error:", error);
        } finally {
            setIsLoading(false);
        }
    };

    const getStatusColor = (status: string) => {
        switch (status) {
            case 'pending': return 'text-gray-500';
            case 'practiced': return 'text-blue-400';
            case 'written': return 'text-yellow-500';
            case 'printed': return 'text-purple-400';
            case 'submitted': return 'text-green-500';
            case 'signed': return 'text-green-600';
            case 'returned': return 'text-red-500';
            default: return 'text-gray-500';
        }
    };

    return (
        <div className="space-y-6">
            <Toaster position="top-center" richColors />

            <Card className="bg-gray-800 border-gray-700">
                <CardHeader>
                    <CardTitle className="text-white">Create New Lab Record</CardTitle>
                </CardHeader>
                <CardContent>
                    <form onSubmit={handleSubmitNewLabRecord} className="grid grid-cols-1 md:grid-cols-2 gap-4">
                        <Input name="experimentNumber" type="number" placeholder="Experiment Number" value={newLabRecord.experimentNumber || ''} onChange={handleInputChange} required className="bg-gray-700 border-gray-600 text-white" />
                        <Input name="title" placeholder="Record Title" value={newLabRecord.title} onChange={handleInputChange} required className="bg-gray-700 border-gray-600 text-white" />
                        
                        <div>
                            <label className="block text-gray-300 text-sm font-bold mb-2">Subject (Optional)</label>
                            <Select name="subjectId" value={newLabRecord.subjectId || ''} onValueChange={(val) => handleSelectChange('subjectId', val)}>
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
                        <Input name="labDate" type="date" placeholder="Lab Date (Optional)" value={newLabRecord.labDate} onChange={handleInputChange} className="bg-gray-700 border-gray-600 text-white" />
                        
                        <Textarea name="aim" placeholder="Aim (Optional)" value={newLabRecord.aim} onChange={handleInputChange} className="bg-gray-700 border-gray-600 text-white col-span-2" />
                        <Textarea name="algorithm" placeholder="Algorithm (Optional)" value={newLabRecord.algorithm} onChange={handleInputChange} className="bg-gray-700 border-gray-600 text-white col-span-2" />
                        <Textarea name="code" placeholder="Code (Optional)" value={newLabRecord.code} onChange={handleInputChange} className="bg-gray-700 border-gray-600 text-white col-span-2" />
                        <Textarea name="output" placeholder="Output (Optional)" value={newLabRecord.output} onChange={handleInputChange} className="bg-gray-700 border-gray-600 text-white col-span-2" />
                        <Textarea name="observations" placeholder="Observations (Optional)" value={newLabRecord.observations} onChange={handleInputChange} className="bg-gray-700 border-gray-600 text-white col-span-2" />
                        <Textarea name="result" placeholder="Result (Optional)" value={newLabRecord.result} onChange={handleInputChange} className="bg-gray-700 border-gray-600 text-white col-span-2" />
                        <Input name="vivaQuestions" placeholder="Viva Questions (comma-separated)" value={newLabRecord.vivaQuestions} onChange={handleInputChange} className="bg-gray-700 border-gray-600 text-white col-span-2" />

                        <div>
                            <label className="block text-gray-300 text-sm font-bold mb-2">Status</label>
                            <Select name="status" value={newLabRecord.status} onValueChange={(val) => handleSelectChange('status', val)} required>
                                <SelectTrigger className="w-full bg-gray-700 border-gray-600 text-white">
                                    <SelectValue placeholder="Select status" />
                                </SelectTrigger>
                                <SelectContent className="bg-gray-700 border-gray-600 text-white">
                                    <SelectItem value="pending">Pending</SelectItem>
                                    <SelectItem value="practiced">Practiced</SelectItem>
                                    <SelectItem value="written">Written</SelectItem>
                                    <SelectItem value="printed">Printed</SelectItem>
                                    <SelectItem value="submitted">Submitted</SelectItem>
                                    <SelectItem value="signed">Signed</SelectItem>
                                    <SelectItem value="returned">Returned</SelectItem>
                                </SelectContent>
                            </Select>
                        </div>
                        
                        <div className="flex items-center space-x-2 md:col-span-2">
                            <input
                                type="checkbox"
                                name="printRequired"
                                checked={newLabRecord.printRequired}
                                onChange={handleInputChange}
                                className="h-4 w-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
                            />
                            <label className="text-gray-300">Print Required?</label>
                        </div>

                        <Input name="pagesToPrint" type="number" placeholder="Pages to Print (Optional)" value={newLabRecord.pagesToPrint || ''} onChange={handleInputChange} className="bg-gray-700 border-gray-600 text-white" />
                        <Input name="marks" type="number" placeholder="Marks (Optional)" value={newLabRecord.marks || ''} onChange={handleInputChange} className="bg-gray-700 border-gray-600 text-white" />
                        <Textarea name="staffRemarks" placeholder="Staff Remarks (Optional)" value={newLabRecord.staffRemarks} onChange={handleInputChange} className="bg-gray-700 border-gray-600 text-white col-span-2" />
                        
                        <div className="md:col-span-2">
                            <Button type="submit" disabled={isLoading} className="w-full bg-blue-600 hover:bg-blue-700">
                                {isLoading ? "Creating..." : "Add Lab Record"}
                            </Button>
                        </div>
                    </form>
                </CardContent>
            </Card>

            <Card className="bg-gray-800 border-gray-700">
                <CardHeader>
                    <CardTitle className="text-white">Your Lab Records</CardTitle>
                </CardHeader>
                <CardContent>
                    {labRecords.length === 0 ? (
                        <p className="text-gray-400">No lab records found. Create one above!</p>
                    ) : (
                        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                            {labRecords.map(record => (
                                <Card key={record.id} className="bg-gray-700 border-gray-600 p-4">
                                    <h3 className="font-bold text-lg text-white">Exp {record.experimentNumber}: {record.title}</h3>
                                    <p className="text-gray-300 text-sm">Subject: {subjects.find(sub => sub.id === record.subjectId)?.name}</p>
                                    {record.labDate && <p className="text-gray-300 text-sm">Lab Date: {new Date(record.labDate).toLocaleDateString()}</p>}
                                    <p className={cn("text-sm font-semibold mt-2", getStatusColor(record.status))}>Status: {record.status}</p>
                                </Card>
                            ))}
                        </div>
                    )}
                </CardContent>
            </Card>
        </div>
    );
}
