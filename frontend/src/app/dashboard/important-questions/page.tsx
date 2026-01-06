"use client";

import { useEffect, useState } from 'react';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { Textarea } from '@/components/ui/textarea';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { Toaster, toast } from 'sonner';
import { useAuthStore } from '@/stores/auth-store';
import { cn } from '@/lib/utils';
import Link from 'next/link';

interface Subject {
    id: string;
    name: string;
    code: string;
}

interface Exam {
    id: string;
    title: string;
}

interface ImportantQuestion {
    id: string;
    questionText: string;
    answerText?: string;
    source?: string;
    unit?: string;
    topic?: string;
    marks?: number;
    isPracticed: boolean;
    subjectId?: string;
    examId?: string;
    tags?: string[];
}

export default function ImportantQuestionsPage() {
    const { token } = useAuthStore();
    const [questions, setQuestions] = useState<ImportantQuestion[]>([]);
    const [subjects, setSubjects] = useState<Subject[]>([]);
    const [exams, setExams] = useState<Exam[]>([]);
    const [newQuestion, setNewQuestion] = useState({
        questionText: '',
        answerText: '',
        source: '',
        unit: '',
        topic: '',
        marks: 0,
        subjectId: '',
        examId: '',
        tags: '',
    });
    const [isLoading, setIsLoading] = useState(false);

    useEffect(() => {
        if (token) {
            fetchQuestions();
            fetchDependencies();
        }
    }, [token]);

    const fetchQuestions = async () => {
        try {
            const res = await fetch("/api/important-questions/subject", { // Fetch by subject for now
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            });
            if (res.ok) {
                const data = await res.json();
                setQuestions(data);
            } else {
                toast.error("Failed to fetch important questions.");
            }
        } catch (error) {
            toast.error("An unexpected error occurred while fetching questions.");
            console.error("Fetch questions error:", error);
        }
    };

    const fetchDependencies = async () => {
        try {
            const [subjectsRes, examsRes] = await Promise.all([
                fetch("/api/subjects", { headers: { Authorization: `Bearer ${token}` } }),
                fetch("/api/exams", { headers: { Authorization: `Bearer ${token}` } }),
            ]);

            if (subjectsRes.ok) setSubjects(await subjectsRes.json());
            else toast.error("Failed to fetch subjects.");
            if (examsRes.ok) setExams(await examsRes.json());
            else toast.error("Failed to fetch exams.");

        } catch (error) {
            toast.error("Error fetching question dependencies.");
            console.error("Fetch dependencies error:", error);
        }
    };

    const handleInputChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        const { name, value, type } = e.target;
        setNewQuestion(prev => ({
            ...prev,
            [name]: type === 'number' ? parseInt(value) : value,
        }));
    };

    const handleSelectChange = (name: string, value: string) => {
        setNewQuestion(prev => ({ ...prev, [name]: value }));
    };

    const handleSubmitNewQuestion = async (e: React.FormEvent) => {
        e.preventDefault();
        setIsLoading(true);

        const payload = {
            ...newQuestion,
            subjectId: newQuestion.subjectId || null,
            examId: newQuestion.examId || null,
            answerText: newQuestion.answerText || null,
            source: newQuestion.source || null,
            unit: newQuestion.unit || null,
            topic: newQuestion.topic || null,
            marks: newQuestion.marks || null,
            tags: newQuestion.tags ? newQuestion.tags.split(',').map(s => s.trim()) : [],
        };

        try {
            const res = await fetch("/api/important-questions", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: `Bearer ${token}`,
                },
                body: JSON.stringify(payload),
            });

            if (res.ok) {
                toast.success("Important question created successfully!");
                setNewQuestion({ // Reset form
                    questionText: '',
                    answerText: '',
                    source: '',
                    unit: '',
                    topic: '',
                    marks: 0,
                    subjectId: '',
                    examId: '',
                    tags: '',
                });
                fetchQuestions(); // Refresh list
            } else {
                const errorData = await res.json();
                toast.error(errorData.error || "Failed to create question.");
            }
        } catch (error) {
            toast.error("An unexpected error occurred while creating question.");
            console.error("Create question error:", error);
        } finally {
            setIsLoading(false);
        }
    };
    
    return (
        <div className="space-y-6">
            <Toaster position="top-center" richColors />

            <Card className="bg-gray-800 border-gray-700">
                <CardHeader>
                    <CardTitle className="text-white">Create New Important Question</CardTitle>
                </CardHeader>
                <CardContent>
                    <form onSubmit={handleSubmitNewQuestion} className="grid grid-cols-1 md:grid-cols-2 gap-4">
                        <Textarea name="questionText" placeholder="Question Text" value={newQuestion.questionText} onChange={handleInputChange} required className="bg-gray-700 border-gray-600 text-white col-span-2" />
                        <Textarea name="answerText" placeholder="Answer Text (Optional)" value={newQuestion.answerText} onChange={handleInputChange} className="bg-gray-700 border-gray-600 text-white col-span-2" />
                        
                        <div>
                            <label className="block text-gray-300 text-sm font-bold mb-2">Subject (Optional)</label>
                            <Select name="subjectId" value={newQuestion.subjectId || ''} onValueChange={(val) => handleSelectChange('subjectId', val)}>
                                <SelectTrigger className="w-full bg-gray-700 border-gray-600 text-white">
                                    <SelectValue placeholder="Select subject" />
                                </SelectTrigger>
                                <SelectContent className="bg-gray-700 border-gray-600 text-white">
                                    {subjects.map(sub => (
                                        <SelectItem key={sub.id} value={sub.id}>{sub.name} ({sub.code})</SelectItem>
                                    ))}
                                </SelectContent>
                            </Select>
                        </div>
                        <div>
                            <label className="block text-gray-300 text-sm font-bold mb-2">Exam (Optional)</label>
                            <Select name="examId" value={newQuestion.examId || ''} onValueChange={(val) => handleSelectChange('examId', val)}>
                                <SelectTrigger className="w-full bg-gray-700 border-gray-600 text-white">
                                    <SelectValue placeholder="Select exam" />
                                </SelectTrigger>
                                <SelectContent className="bg-gray-700 border-gray-600 text-white">
                                    {exams.map(exam => (
                                        <SelectItem key={exam.id} value={exam.id}>{exam.title}</SelectItem>
                                    ))}
                                </SelectContent>
                            </Select>
                        </div>

                        <Input name="source" placeholder="Source (e.g., PYQ, Staff Notes)" value={newQuestion.source} onChange={handleInputChange} className="bg-gray-700 border-gray-600 text-white" />
                        <Input name="unit" placeholder="Unit (Optional)" value={newQuestion.unit} onChange={handleInputChange} className="bg-gray-700 border-gray-600 text-white" />
                        <Input name="topic" placeholder="Topic (Optional)" value={newQuestion.topic} onChange={handleInputChange} className="bg-gray-700 border-gray-600 text-white" />
                        <Input name="marks" type="number" placeholder="Marks (Optional)" value={newQuestion.marks || ''} onChange={handleInputChange} className="bg-gray-700 border-gray-600 text-white" />
                        <Input name="tags" placeholder="Tags (comma-separated)" value={newQuestion.tags} onChange={handleInputChange} className="bg-gray-700 border-gray-600 text-white col-span-2" />
                        
                        <div className="md:col-span-2">
                            <Button type="submit" disabled={isLoading} className="w-full bg-blue-600 hover:bg-blue-700">
                                {isLoading ? "Creating..." : "Add Question"}
                            </Button>
                        </div>
                    </form>
                </CardContent>
            </Card>

            <Card className="bg-gray-800 border-gray-700">
                <CardHeader>
                    <CardTitle className="text-white">Your Important Questions</CardTitle>
                </CardHeader>
                <CardContent>
                    {questions.length === 0 ? (
                        <p className="text-gray-400">No important questions found. Create one above!</p>
                    ) : (
                        <div className="space-y-4">
                            {questions.map(q => (
                                <Card key={q.id} className="bg-gray-700 border-gray-600 p-4">
                                    <h3 className="font-bold text-lg text-white">{q.questionText}</h3>
                                    {q.answerText && <p className="text-gray-300 text-sm mt-1">Answer: {q.answerText}</p>}
                                    {q.subjectId && <p className="text-gray-300 text-sm">Subject: {subjects.find(sub => sub.id === q.subjectId)?.name}</p>}
                                    {q.examId && <p className="text-gray-300 text-sm">Exam: {exams.find(exam => exam.id === q.examId)?.title}</p>}
                                    {q.unit && <p className="text-gray-300 text-sm">Unit: {q.unit}</p>}
                                    {q.topic && <p className="text-gray-300 text-sm">Topic: {q.topic}</p>}
                                    {q.tags && q.tags.length > 0 && <p className="text-gray-300 text-sm">Tags: {q.tags.join(', ')}</p>}
                                </Card>
                            ))}
                        </div>
                    )}
                </CardContent>
            </Card>
        </div>
    );
}
