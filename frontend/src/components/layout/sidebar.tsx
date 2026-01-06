"use client";

import Link from 'next/link';
import { usePathname } from 'next/navigation';
import { cn } from '@/lib/utils'; // For conditional class names

const navItems = [
    { name: 'Dashboard', href: '/dashboard' },
    { name: 'Assignments', href: '/dashboard/assignments' },
    { name: 'Exams', href: '/dashboard/exams' },
    { name: 'Important Questions', href: '/dashboard/important-questions' },
    { name: 'Lab Records', href: '/dashboard/lab-records' },
    { name: 'Documents', href: '/dashboard/documents' },
    { name: 'Subjects', href: '/dashboard/subjects' },
    { name: 'Staff', href: '/dashboard/staff' },
    { name: 'Venues', href: '/dashboard/venues' },
    { name: 'Timetable', href: '/dashboard/timetable' },
];

export default function Sidebar() {
    const pathname = usePathname();

    return (
        <aside className="w-64 bg-gray-800 border-r border-gray-700 p-4 flex flex-col">
            <div className="text-2xl font-bold text-white mb-8">Campus Pilot</div>
            <nav className="flex-1">
                <ul className="space-y-2">
                    {navItems.map((item) => (
                        <li key={item.name}>
                            <Link href={item.href}>
                                <div className={cn(
                                    "flex items-center p-2 text-gray-300 rounded-md hover:bg-gray-700",
                                    pathname === item.href && "bg-gray-700 text-white"
                                )}>
                                    {item.name}
                                </div>
                            </Link>
                        </li>
                    ))}
                </ul>
            </nav>
        </aside>
    );
}
