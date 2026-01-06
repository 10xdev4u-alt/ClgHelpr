"use client";

import { useAuthStore } from "@/stores/auth-store";
import { Button } from "@/components/ui/button";
import { useRouter } from "next/navigation";

export default function DashboardPage() {
    const { token, setToken } = useAuthStore();
    const router = useRouter();

    const handleLogout = () => {
        setToken(null);
        router.push("/login");
    };

    return (
        <div className="flex flex-col items-center justify-center min-h-screen bg-gray-900 text-white">
            <div className="p-8 bg-gray-800 rounded-lg shadow-lg text-center">
                <h1 className="text-3xl font-bold mb-4">Welcome to the Dashboard!</h1>
                <p className="mb-6 text-gray-400">You are successfully logged in.</p>
                <div className="mb-6 p-4 bg-gray-700 rounded-md break-all">
                    <p className="font-mono text-sm text-green-400">
                        <span className="font-bold text-gray-300">Your JWT:</span> {token}
                    </p>
                </div>
                <Button onClick={handleLogout} variant="destructive">
                    Logout
                </Button>
            </div>
        </div>
    );
}
