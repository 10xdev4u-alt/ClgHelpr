"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Toaster, toast } from "sonner";

export default function RegisterPage() {
    const router = useRouter();
    const [formData, setFormData] = useState({
        fullName: "",
        email: "",
        password: "",
        registerNumber: "",
        department: "CSE",
        year: 3,
        semester: 6,
    });
    const [isLoading, setIsLoading] = useState(false);

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value, type } = e.target;
        setFormData((prev) => ({
            ...prev,
            [name]: type === 'number' ? parseInt(value, 10) : value,
        }));
    };

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        setIsLoading(true);

        try {
            const res = await fetch("/api/auth/register", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify(formData),
            });

            if (res.ok) {
                toast.success("Registration successful! Redirecting to login...");
                setTimeout(() => router.push("/login"), 2000);
            } else {
                const errorData = await res.json();
                toast.error(errorData.error || "Registration failed. Please try again.");
            }
        } catch (error) {
            toast.error("An unexpected error occurred. Please check your connection.");
            console.error(error);
        } finally {
            setIsLoading(false);
        }
    };

    return (
        <>
            <Toaster position="top-center" richColors />
            <div className="flex items-center justify-center min-h-screen bg-gray-900 text-white">
                <Card className="w-full max-w-md bg-gray-800 border-gray-700">
                    <CardHeader className="text-center">
                        <CardTitle className="text-2xl font-bold">Create an Account</CardTitle>
                        <CardDescription className="text-gray-400">
                            Join Campus Pilot to streamline your college life.
                        </CardDescription>
                    </CardHeader>
                    <CardContent>
                        <form onSubmit={handleSubmit} className="space-y-4">
                            <Input name="fullName" placeholder="Full Name" value={formData.fullName} onChange={handleChange} required className="bg-gray-700 border-gray-600 text-white" />
                            <Input name="email" type="email" placeholder="Email" value={formData.email} onChange={handleChange} required className="bg-gray-700 border-gray-600 text-white" />
                            <Input name="password" type="password" placeholder="Password" value={formData.password} onChange={handleChange} required className="bg-gray-700 border-gray-600 text-white" />
                            <Input name="registerNumber" placeholder="Register Number" value={formData.registerNumber} onChange={handleChange} required className="bg-gray-700 border-gray-600 text-white" />
                            <div className="grid grid-cols-2 gap-4">
                                <Input name="department" placeholder="Department" value={formData.department} onChange={handleChange} required className="bg-gray-700 border-gray-600 text-white" />
                                <Input name="year" type="number" placeholder="Year" value={formData.year} onChange={handleChange} required className="bg-gray-700 border-gray-600 text-white" />
                            </div>
                            <Input name="semester" type="number" placeholder="Semester" value={formData.semester} onChange={handleChange} required className="bg-gray-700 border-gray-600 text-white" />
                            <Button type="submit" className="w-full bg-blue-600 hover:bg-blue-700" disabled={isLoading}>
                                {isLoading ? "Registering..." : "Create Account"}
                            </Button>
                        </form>
                    </CardContent>
                </Card>
            </div>
        </>
    );
}
