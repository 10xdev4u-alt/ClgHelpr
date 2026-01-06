"use client";

import { useAuthStore } from "@/stores/auth-store";
import { Button } from "@/components/ui/button";
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";

interface User {
    fullName: string;
    email: string;
}

export default function DashboardPage() {
    const { token, setToken } = useAuthStore();
    const router = useRouter();
    const [user, setUser] = useState<User | null>(null);

    useEffect(() => {
        const fetchUser = async () => {
            if (token) {
                try {
                    const res = await fetch("/api/me", {
                        headers: {
                            Authorization: `Bearer ${token}`,
                        },
                    });
                    if (res.ok) {
                        const userData = await res.json();
                        setUser(userData);
                    } else {
                        // Token might be expired/invalid
                        setToken(null);
                        router.push("/login");
                    }
                } catch (error) {
                    console.error("Failed to fetch user profile", error);
                    setToken(null);
                    router.push("/login");
                }
            }
        };
        fetchUser();
    }, [token, router, setToken]);

    const handleLogout = () => {
        setToken(null);
        router.push("/login");
    };
    
    if (!user) {
        return <div className="flex items-center justify-center min-h-screen bg-gray-900 text-white">Loading profile...</div>;
    }

    return (
        <div className="flex flex-col items-center justify-center min-h-screen bg-gray-900 text-white">
            <div className="p-8 bg-gray-800 rounded-lg shadow-lg text-center">
                <h1 className="text-3xl font-bold mb-4">Welcome, {user.fullName}!</h1>
                <p className="mb-6 text-gray-400">You are successfully logged in. Your email is {user.email}.</p>
                <Button onClick={handleLogout} variant="destructive">
                    Logout
                </Button>
            </div>
        </div>
    );
}
