"use client";

import { useEffect, useState } from 'react';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { Toaster, toast } from 'sonner';
import { useAuthStore } from '@/stores/auth-store';

interface Staff {
    id: string;
    name: string;
    title?: string;
    email?: string;
    phone?: string;
    department?: string;
    designation?: string;
    cabin?: string;
}

export default function StaffPage() {
    const { token } = useAuthStore();
    const [staffMembers, setStaffMembers] = useState<Staff[]>([]);
    const [newStaff, setNewStaff] = useState<Omit<Staff, 'id' | 'createdAt'>>({
        name: '',
        department: 'CSE',
    });
    const [isLoading, setIsLoading] = useState(false);

    useEffect(() => {
        if (token) {
            fetchStaff();
        }
    }, [token]);

    const fetchStaff = async () => {
        try {
            const res = await fetch("/api/staff", {
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            });
            if (res.ok) {
                const data = await res.json();
                setStaffMembers(data);
            } else {
                toast.error("Failed to fetch staff members.");
            }
        } catch (error) {
            toast.error("An unexpected error occurred while fetching staff members.");
            console.error("Fetch staff error:", error);
        }
    };

    const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target;
        setNewStaff(prev => ({ ...prev, [name]: value }));
    };

    const handleSubmitNewStaff = async (e: React.FormEvent) => {
        e.preventDefault();
        setIsLoading(true);

        try {
            const res = await fetch("/api/timetable/staff", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: `Bearer ${token}`,
                },
                body: JSON.stringify(newStaff),
            });

            if (res.ok) {
                toast.success("Staff member created successfully!");
                setNewStaff({ name: '', department: 'CSE' }); // Reset form
                fetchStaff(); // Refresh list
            } else {
                const errorData = await res.json();
                toast.error(errorData.error || "Failed to create staff member.");
            }
        } catch (error) {
            toast.error("An unexpected error occurred while creating staff member.");
            console.error("Create staff error:", error);
        } finally {
            setIsLoading(false);
        }
    };

    return (
        <div className="space-y-6">
            <Toaster position="top-center" richColors />

            <Card className="bg-gray-800 border-gray-700">
                <CardHeader>
                    <CardTitle className="text-white">Create New Staff Member</CardTitle>
                </CardHeader>
                <CardContent>
                    <form onSubmit={handleSubmitNewStaff} className="grid grid-cols-1 md:grid-cols-2 gap-4">
                        <Input name="name" placeholder="Name" value={newStaff.name} onChange={handleInputChange} required className="bg-gray-700 border-gray-600 text-white" />
                        <Input name="title" placeholder="Title (e.g., Dr., Ms.)" value={newStaff.title || ''} onChange={handleInputChange} className="bg-gray-700 border-gray-600 text-white" />
                        <Input name="email" type="email" placeholder="Email (Optional)" value={newStaff.email || ''} onChange={handleInputChange} className="bg-gray-700 border-gray-600 text-white" />
                        <Input name="phone" placeholder="Phone (Optional)" value={newStaff.phone || ''} onChange={handleInputChange} className="bg-gray-700 border-gray-600 text-white" />
                        <Input name="department" placeholder="Department" value={newStaff.department} onChange={handleInputChange} required className="bg-gray-700 border-gray-600 text-white" />
                        <Input name="designation" placeholder="Designation (Optional)" value={newStaff.designation || ''} onChange={handleInputChange} className="bg-gray-700 border-gray-600 text-white" />
                        <Input name="cabin" placeholder="Cabin (Optional)" value={newStaff.cabin || ''} onChange={handleInputChange} className="bg-gray-700 border-gray-600 text-white" />
                        <div className="md:col-span-2">
                            <Button type="submit" disabled={isLoading} className="w-full bg-blue-600 hover:bg-blue-700">
                                {isLoading ? "Creating..." : "Add Staff Member"}
                            </Button>
                        </div>
                    </form>
                </CardContent>
            </Card>

            <Card className="bg-gray-800 border-gray-700">
                <CardHeader>
                    <CardTitle className="text-white">Staff Members List</CardTitle>
                </CardHeader>
                <CardContent>
                    {staffMembers.length === 0 ? (
                        <p className="text-gray-400">No staff members found. Create one above!</p>
                    ) : (
                        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                            {staffMembers.map(staff => (
                                <Card key={staff.id} className="bg-gray-700 border-gray-600 p-4">
                                    <h3 className="font-bold text-lg text-white">{staff.name} {staff.title && `(${staff.title})`}</h3>
                                    <p className="text-gray-300 text-sm">Dept: {staff.department}</p>
                                    {staff.designation && <p className="text-gray-300 text-sm">Desig: {staff.designation}</p>}
                                    {staff.email && <p className="text-gray-300 text-sm">Email: {staff.email}</p>}
                                </Card>
                            ))}
                        </div>
                    )}
                </CardContent>
            </Card>
        </div>
    );
}
