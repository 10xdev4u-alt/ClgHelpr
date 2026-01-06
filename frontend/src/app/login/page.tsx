"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import Link from "next/link";
import { Toaster, toast } from "sonner";
import { useAuthStore } from "@/stores/auth-store";

export default function LoginPage() {
    const router = useRouter();
    const { setToken } = useAuthStore();
    const [formData, setFormData] = useState({ email: "", password: "" });
    const [isLoading, setIsLoading] = useState(false);

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setFormData((prev) => ({ ...prev, [e.target.name]: e.target.value }));
    };

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        setIsLoading(true);

        try {
            const res = await fetch("/api/auth/login", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify(formData),
            });

            if (res.ok) {
                const { token } = await res.json();
                setToken(token);
                toast.success("Login successful! Redirecting to dashboard...");
                setTimeout(() => router.push("/dashboard"), 1000);
            } else {
                const errorData = await res.json();
                toast.error(errorData.error || "Login failed. Please check your credentials.");
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
                        <CardTitle className="text-2xl font-bold">Login to Campus Pilot</CardTitle>
                        <CardDescription className="text-gray-400">
                            Welcome back! Please enter your credentials.
                        </CardDescription>
                    </CardHeader>
                    <CardContent>
                        <form onSubmit={handleSubmit} className="space-y-4">
                            <Input name="email" type="email" placeholder="Email" value={formData.email} onChange={handleChange} required className="bg-gray-700 border-gray-600 text-white" />
                            <Input name="password" type="password" placeholder="Password" value={formData.password} onChange={handleChange} required className="bg-gray-700 border-gray-600 text-white" />
                            <Button type="submit" className="w-full bg-blue-600 hover:bg-blue-700" disabled={isLoading}>
                                {isLoading ? "Logging in..." : "Login"}
                            </Button>
                        </form>
                        <p className="text-center text-sm text-gray-400 mt-4">
                            Don't have an account?{" "}
                            <Link href="/register" className="text-blue-400 hover:underline">
                                Register here
                            </Link>
                        </p>
                    </CardContent>
                </Card>
            </div>
        </>
    );
}
