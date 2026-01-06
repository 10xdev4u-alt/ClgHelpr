"use client";

import { useEffect, useState } from 'react';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { Toaster, toast } from 'sonner';
import { useAuthStore } from '@/stores/auth-store';

interface Subject {
    id: string;
    name: string;
    code: string;
}

interface Staff {
    id: string;
    name: string;
}

interface Venue {
    id: string;
    name: string;
}

interface TimetableSlot {
    id: string;
    userId: string;
    subjectId: string;
    staffId?: string;
    venueId?: string;
    dayOfWeek: number; // 0-6, Sunday-Saturday
    startTime: string; // HH:MM
    endTime: string;   // HH:MM
    periodNumber?: number;
    slotType: string;
    isRecurring: boolean;
    specificDate?: string; // YYYY-MM-DD
    notes?: string;
    batchFilter?: string;
}

export default function TimetablePage() {
    const { token } = useAuthStore();
    const [subjects, setSubjects] = useState<Subject[]>([]);
    const [staffMembers, setStaffMembers] = useState<Staff[]>([]);
    const [venues, setVenues] = useState<Venue[]>([]);
    const [timetable, setTimetable] = useState<TimetableSlot[]>([]);
    const [isLoading, setIsLoading] = useState(false);

    const [newSlot, setNewSlot] = useState<Omit<TimetableSlot, 'id' | 'userId'>>({
        subjectId: '',
        dayOfWeek: 1, // Default to Monday
        startTime: '08:30',
        endTime: '09:20',
        slotType: 'lecture',
        isRecurring: true,
    });

    useEffect(() => {
        if (token) {
            fetchDependencies();
            fetchTimetable();
        }
    }, [token]);

    const fetchDependencies = async () => {
        try {
            const [subjectsRes, staffRes, venuesRes] = await Promise.all([
                fetch("/api/subjects", { headers: { Authorization: `Bearer ${token}` } }),
                fetch("/api/staff", { headers: { Authorization: `Bearer ${token}` } }),
                fetch("/api/venues", { headers: { Authorization: `Bearer ${token}` } }),
            ]);

            if (subjectsRes.ok) setSubjects(await subjectsRes.json());
            else toast.error("Failed to fetch subjects.");
            if (staffRes.ok) setStaffMembers(await staffRes.json());
            else toast.error("Failed to fetch staff.");
            if (venuesRes.ok) setVenues(await venuesRes.json());
            else toast.error("Failed to fetch venues.");

        } catch (error) {
            toast.error("Error fetching timetable dependencies.");
            console.error("Fetch dependencies error:", error);
        }
    };

    const fetchTimetable = async () => {
        try {
            // For now, fetch for today (Monday = 1)
            const today = new Date();
            const dayOfWeek = today.getDay(); // 0 for Sunday, 1 for Monday
            
            const res = await fetch(`/api/timetable/day/${dayOfWeek}`, {
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            });
            if (res.ok) {
                const data = await res.json();
                setTimetable(data);
            } else {
                toast.error("Failed to fetch timetable.");
            }
        } catch (error) {
            toast.error("An unexpected error occurred while fetching timetable.");
            console.error("Fetch timetable error:", error);
        }
    };

    const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value, type, checked } = e.target;
        setNewSlot(prev => ({
            ...prev,
            [name]: type === 'checkbox' ? checked : (type === 'number' ? parseInt(value) : value),
        }));
    };

    const handleSelectChange = (name: string, value: string) => {
        setNewSlot(prev => ({ ...prev, [name]: value }));
    };

    const handleSubmitNewSlot = async (e: React.FormEvent) => {
        e.preventDefault();
        setIsLoading(true);

        try {
            const res = await fetch("/api/timetable/slots", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: `Bearer ${token}`,
                },
                body: JSON.stringify({
                    ...newSlot,
                    dayOfWeek: parseInt(newSlot.dayOfWeek as any), // Ensure number
                    periodNumber: newSlot.periodNumber ? parseInt(newSlot.periodNumber as any) : null,
                    // Specific date needs parsing if not recurring
                }),
            });

            if (res.ok) {
                toast.success("Timetable slot created successfully!");
                // Reset form (consider a more complete reset)
                setNewSlot({
                    subjectId: '',
                    dayOfWeek: 1,
                    startTime: '08:30',
                    endTime: '09:20',
                    slotType: 'lecture',
                    isRecurring: true,
                });
                fetchTimetable(); // Refresh list
            } else {
                const errorData = await res.json();
                toast.error(errorData.error || "Failed to create timetable slot.");
            }
        } catch (error) {
            toast.error("An unexpected error occurred while creating slot.");
            console.error("Create slot error:", error);
        } finally {
            setIsLoading(false);
        }
    };

    const daysOfWeek = [
        "Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"
    ];

    return (
        <div className="space-y-6">
            <Toaster position="top-center" richColors />

            <Card className="bg-gray-800 border-gray-700">
                <CardHeader>
                    <CardTitle className="text-white">Create New Timetable Slot</CardTitle>
                </CardHeader>
                <CardContent>
                    <form onSubmit={handleSubmitNewSlot} className="grid grid-cols-1 md:grid-cols-2 gap-4">
                        <div>
                            <label className="block text-gray-300 text-sm font-bold mb-2">Subject</label>
                            <Select name="subjectId" value={newSlot.subjectId} onValueChange={(val) => handleSelectChange('subjectId', val)} required>
                                <SelectTrigger className="w-full bg-gray-700 border-gray-600 text-white">
                                    <SelectValue placeholder="Select a subject" />
                                </SelectTrigger>
                                <SelectContent className="bg-gray-700 border-gray-600 text-white">
                                    {subjects.map(sub => (
                                        <SelectItem key={sub.id} value={sub.id}>{sub.name} ({sub.code})</SelectItem>
                                    ))}
                                </SelectContent>
                            </Select>
                        </div>

                        <div>
                            <label className="block text-gray-300 text-sm font-bold mb-2">Day of Week</label>
                            <Select name="dayOfWeek" value={newSlot.dayOfWeek.toString()} onValueChange={(val) => handleSelectChange('dayOfWeek', parseInt(val).toString())} required>
                                <SelectTrigger className="w-full bg-gray-700 border-gray-600 text-white">
                                    <SelectValue placeholder="Select day" />
                                </SelectTrigger>
                                <SelectContent className="bg-gray-700 border-gray-600 text-white">
                                    {daysOfWeek.map((day, index) => (
                                        <SelectItem key={index} value={index.toString()}>{day}</SelectItem>
                                    ))}
                                </SelectContent>
                            </Select>
                        </div>
                        
                        <Input name="startTime" type="time" label="Start Time" value={newSlot.startTime} onChange={handleInputChange} required className="bg-gray-700 border-gray-600 text-white" />
                        <Input name="endTime" type="time" label="End Time" value={newSlot.endTime} onChange={handleInputChange} required className="bg-gray-700 border-gray-600 text-white" />

                        <div>
                            <label className="block text-gray-300 text-sm font-bold mb-2">Staff (Optional)</label>
                            <Select name="staffId" value={newSlot.staffId || ''} onValueChange={(val) => handleSelectChange('staffId', val)}>
                                <SelectTrigger className="w-full bg-gray-700 border-gray-600 text-white">
                                    <SelectValue placeholder="Select staff" />
                                </SelectTrigger>
                                <SelectContent className="bg-gray-700 border-gray-600 text-white">
                                    {staffMembers.map(staff => (
                                        <SelectItem key={staff.id} value={staff.id}>{staff.name}</SelectItem>
                                    ))}
                                </SelectContent>
                            </Select>
                        </div>
                        <div>
                            <label className="block text-gray-300 text-sm font-bold mb-2">Venue (Optional)</label>
                            <Select name="venueId" value={newSlot.venueId || ''} onValueChange={(val) => handleSelectChange('venueId', val)}>
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

                        <Input name="slotType" placeholder="Slot Type (e.g., lecture, lab)" value={newSlot.slotType} onChange={handleInputChange} required className="bg-gray-700 border-gray-600 text-white" />
                        <Input name="periodNumber" type="number" placeholder="Period Number (Optional)" value={newSlot.periodNumber || ''} onChange={handleInputChange} className="bg-gray-700 border-gray-600 text-white" />
                        
                        <div className="md:col-span-2 flex items-center space-x-2">
                            <input
                                type="checkbox"
                                name="isRecurring"
                                checked={newSlot.isRecurring}
                                onChange={handleInputChange}
                                className="h-4 w-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
                            />
                            <label className="text-gray-300">Is Recurring?</label>
                        </div>

                        {/* Add specificDate input if !isRecurring later */}

                        <div className="md:col-span-2">
                            <Button type="submit" disabled={isLoading} className="w-full bg-blue-600 hover:bg-blue-700">
                                {isLoading ? "Creating..." : "Add Timetable Slot"}
                            </Button>
                        </div>
                    </form>
                </CardContent>
            </Card>

            <Card className="bg-gray-800 border-gray-700">
                <CardHeader>
                    <CardTitle className="text-white">Today's Timetable</CardTitle>
                </CardHeader>
                <CardContent>
                    {timetable.length === 0 ? (
                        <p className="text-gray-400">No timetable slots for today. Add some above!</p>
                    ) : (
                        <div className="space-y-4">
                            {timetable.map(slot => (
                                <Card key={slot.id} className="bg-gray-700 border-gray-600 p-4">
                                    <h3 className="font-bold text-lg text-white">
                                        {slot.startTime} - {slot.endTime}: {subjects.find(sub => sub.id === slot.subjectId)?.name || 'Unknown Subject'}
                                    </h3>
                                    <p className="text-gray-300 text-sm">Type: {slot.slotType}</p>
                                    {slot.staffId && <p className="text-gray-300 text-sm">Staff: {staffMembers.find(s => s.id === slot.staffId)?.name}</p>}
                                    {slot.venueId && <p className="text-gray-300 text-sm">Venue: {venues.find(v => v.id === slot.venueId)?.name}</p>}
                                    <p className="text-gray-300 text-sm">Day: {daysOfWeek[slot.dayOfWeek]}</p>
                                </Card>
                            ))}
                        </div>
                    )}
                </CardContent>
            </Card>
        </div>
    );
}
