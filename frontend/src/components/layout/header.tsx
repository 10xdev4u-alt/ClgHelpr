"use client";

import { useAuthStore } from '@/stores/auth-store';
import { useRouter } from 'next/navigation';
import { Button } from '@/components/ui/button';
import { toast } from 'sonner';

export default function Header() {
    const { setToken } = useAuthStore();
    const router = useRouter();

    const handleLogout = () => {
        setToken(null);
        toast.info("Logged out successfully.");
        router.push("/login");
    };

    return (
        <header className="flex items-center justify-between h-16 px-6 bg-gray-800 border-b border-gray-700">
            <h1 className="text-xl font-semibold text-white">Dashboard</h1>
            <Button onClick={handleLogout} variant="destructive" size="sm">
                Logout
            </Button>
        </header>
    );
}
