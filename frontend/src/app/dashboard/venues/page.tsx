"use client";

import { useEffect, useState } from 'react';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { Toaster, toast } from 'sonner';
import { useAuthStore } from '@/stores/auth-store';

interface Venue {
    id: string;
    name: string;
    building?: string;
    floor?: number;
    capacity?: number;
    type: string;
    facilities?: object; // Assuming JSONB maps to object
}

export default function VenuesPage() {
    const { token } = useAuthStore();
    const [venues, setVenues] = useState<Venue[]>([]);
    const [newVenue, setNewVenue] = useState<Omit<Venue, 'id' | 'createdAt'>>({
        name: '',
        type: 'classroom',
    });
    const [isLoading, setIsLoading] = useState(false);

    useEffect(() => {
        if (token) {
            fetchVenues();
        }
    }, [token]);

    const fetchVenues = async () => {
        try {
            const res = await fetch("/api/venues", {
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            });
            if (res.ok) {
                const data = await res.json();
                setVenues(data);
            } else {
                toast.error("Failed to fetch venues.");
            }
        } catch (error) {
            toast.error("An unexpected error occurred while fetching venues.");
            console.error("Fetch venues error:", error);
        }
    };

    const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value, type } = e.target;
        setNewVenue(prev => ({
            ...prev,
            [name]: type === 'number' ? parseInt(value) : value,
        }));
    };

    const handleSubmitNewVenue = async (e: React.FormEvent) => {
        e.preventDefault();
        setIsLoading(true);

        try {
            const res = await fetch("/api/timetable/venues", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: `Bearer ${token}`,
                },
                body: JSON.stringify(newVenue),
            });

            if (res.ok) {
                toast.success("Venue created successfully!");
                setNewVenue({ name: '', type: 'classroom' }); // Reset form
                fetchVenues(); // Refresh list
            } else {
                const errorData = await res.json();
                toast.error(errorData.error || "Failed to create venue.");
            }
        } catch (error) {
            toast.error("An unexpected error occurred while creating venue.");
            console.error("Create venue error:", error);
        } finally {
            setIsLoading(false);
        }
    };

    return (
        <div className="space-y-6">
            <Toaster position="top-center" richColors />

            <Card className="bg-gray-800 border-gray-700">
                <CardHeader>
                    <CardTitle className="text-white">Create New Venue</CardTitle>
                </CardHeader>
                <CardContent>
                    <form onSubmit={handleSubmitNewVenue} className="grid grid-cols-1 md:grid-cols-2 gap-4">
                        <Input name="name" placeholder="Venue Name (e.g., CSE Lab 1)" value={newVenue.name} onChange={handleInputChange} required className="bg-gray-700 border-gray-600 text-white" />
                        <Input name="type" placeholder="Type (e.g., classroom, lab)" value={newVenue.type} onChange={handleInputChange} required className="bg-gray-700 border-gray-600 text-white" />
                        <Input name="building" placeholder="Building (Optional)" value={newVenue.building || ''} onChange={handleInputChange} className="bg-gray-700 border-gray-600 text-white" />
                        <Input name="floor" type="number" placeholder="Floor (Optional)" value={newVenue.floor || ''} onChange={handleInputChange} className="bg-gray-700 border-gray-600 text-white" />
                        <Input name="capacity" type="number" placeholder="Capacity (Optional)" value={newVenue.capacity || ''} onChange={handleInputChange} className="bg-gray-700 border-gray-600 text-white" />
                        {/* Facilities (JSONB) input can be added later as a more complex component */}
                        <div className="md:col-span-2">
                            <Button type="submit" disabled={isLoading} className="w-full bg-blue-600 hover:bg-blue-700">
                                {isLoading ? "Creating..." : "Add Venue"}
                            </Button>
                        </div>
                    </form>
                </CardContent>
            </Card>

            <Card className="bg-gray-800 border-gray-700">
                <CardHeader>
                    <CardTitle className="text-white">Venues List</CardTitle>
                </CardHeader>
                <CardContent>
                    {venues.length === 0 ? (
                        <p className="text-gray-400">No venues found. Create one above!</p>
                    ) : (
                        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                            {venues.map(venue => (
                                <Card key={venue.id} className="bg-gray-700 border-gray-600 p-4">
                                    <h3 className="font-bold text-lg text-white">{venue.name} ({venue.type})</h3>
                                    {venue.building && <p className="text-gray-300 text-sm">Building: {venue.building}</p>}
                                    {venue.floor && <p className="text-gray-300 text-sm">Floor: {venue.floor}</p>}
                                    {venue.capacity && <p className="text-gray-300 text-sm">Capacity: {venue.capacity}</p>}
                                </Card>
                            ))}
                        </div>
                    )}
                </CardContent>
            </Card>
        </div>
    );
}
