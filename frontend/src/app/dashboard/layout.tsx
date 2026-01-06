import AuthGuard from '@/components/shared/auth-guard';
import Sidebar from '@/components/layout/sidebar'; // We'll create this next
import Header from '@/components/layout/header'; // And this

export default function DashboardLayout({
    children,
}: {
    children: React.ReactNode;
}) {
    return (
        <AuthGuard>
            <div className="flex min-h-screen bg-gray-900 text-white">
                <Sidebar />
                <div className="flex flex-col flex-1">
                    <Header />
                    <main className="flex-1 p-6">
                        {children}
                    </main>
                </div>
            </div>
        </AuthGuard>
    );
}
