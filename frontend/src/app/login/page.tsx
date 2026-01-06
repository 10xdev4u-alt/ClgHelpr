"use client";

import { Button } from "@/components/ui/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import Link from "next/link";

export default function LoginPage() {
    return (
        <div className="flex items-center justify-center min-h-screen bg-gray-900 text-white">
            <Card className="w-full max-w-md bg-gray-800 border-gray-700">
                <CardHeader className="text-center">
                    <CardTitle className="text-2xl font-bold">Login to Campus Pilot</CardTitle>
                    <CardDescription className="text-gray-400">
                        Welcome back! Please enter your credentials.
                    </CardDescription>
                </CardHeader>
                <CardContent>
                    <form className="space-y-4">
                        <Input name="email" type="email" placeholder="Email" required className="bg-gray-700 border-gray-600 text-white" />
                        <Input name="password" type="password" placeholder="Password" required className="bg-gray-700 border-gray-600 text-white" />
                        <Button type="submit" className="w-full bg-blue-600 hover:bg-blue-700">
                            Login
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
    );
}
