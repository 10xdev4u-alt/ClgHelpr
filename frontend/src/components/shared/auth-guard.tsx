"use client";

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import { useAuthStore } from '@/stores/auth-store';

export default function AuthGuard({ children }: { children: React.ReactNode }) {
    const router = useRouter();
    const { isLoggedIn } = useAuthStore();
    const [isAuthChecked, setIsAuthChecked] = useState(false);

    useEffect(() => {
        if (!isLoggedIn()) {
            router.replace('/login');
        } else {
            setIsAuthChecked(true);
        }
    }, [isLoggedIn, router]);

    if (!isAuthChecked) {
        // You can render a loading spinner here
        return <div className="flex items-center justify-center min-h-screen bg-gray-900 text-white">Loading...</div>;
    }

    return <>{children}</>;
}
