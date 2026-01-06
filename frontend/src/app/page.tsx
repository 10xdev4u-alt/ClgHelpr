"use client";

import { Button } from "@/components/ui/button";
import { useState } from "react";

export default function Home() {
  const [message, setMessage] = useState("Click the button to fetch from API.");

  const fetchApi = async () => {
    try {
      const res = await fetch("/api");
      const data = await res.json();
      setMessage(data.message || "No message found.");
    } catch (error) {
      setMessage("Failed to fetch from API.");
      console.error(error);
    }
  };

  return (
    <div className="flex flex-col min-h-screen bg-gray-900 text-white">
      <header className="px-4 lg:px-6 h-14 flex items-center bg-gray-800/40 border-b border-gray-700">
        <a className="flex items-center justify-center" href="#">
          <span className="text-xl font-bold">Campus Pilot</span>
        </a>
        <nav className="ml-auto flex gap-4 sm:gap-6">
          <Button variant="outline" className="text-white border-gray-600 hover:bg-gray-700 hover:text-white">
            Login
          </Button>
        </nav>
      </header>
      <main className="flex-1 flex flex-col items-center justify-center p-4 text-center">
        <div className="space-y-4">
          <h1 className="text-4xl md:text-5xl font-bold tracking-tighter">
            Your AI-Powered Academic Co-Pilot
          </h1>
          <p className="max-w-[600px] text-gray-400 md:text-xl">
            Never miss a deadline, lose a note, or feel unprepared again.
            Welcome to the future of college management.
          </p>
          <div className="space-y-2 pt-4">
            <Button onClick={fetchApi}>Fetch Welcome Message</Button>
            <p className="text-sm text-gray-500 italic">{message}</p>
          </div>
        </div>
      </main>
      <footer className="flex items-center justify-center p-4 bg-gray-800/40 border-t border-gray-700">
        <p className="text-sm text-gray-500">
          Built for PrinceTheProgrammer &copy; {new Date().getFullYear()}
        </p>
      </footer>
    </div>
  );
}
