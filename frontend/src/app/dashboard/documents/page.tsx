"use client";

import { useEffect, useState } from 'react';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { Textarea } from '@/components/ui/textarea';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { Checkbox } from '@/components/ui/checkbox'; // Assuming Checkbox component
import { Toaster, toast } from 'sonner';
import { useAuthStore } from '@/stores/auth-store';
import { cn } from '@/lib/utils';
import Link from 'next/link';

interface Subject {
    id: string;
    name: string;
    code: string;
}

interface Document {
    id: string;
    title: string;
    description?: string;
    documentType: string;
    fileName?: string;
    fileType?: string;
    fileSize?: number;
    fileUrl: string;
    storageKey?: string;
    folder?: string;
    tags?: string[];
    isPublic: boolean;
    subjectId?: string;
    // ... other fields as needed for display
}

export default function DocumentsPage() {
    const { token } = useAuthStore();
    const [documents, setDocuments] = useState<Document[]>([]);
    const [subjects, setSubjects] = useState<Subject[]>([]);
    const [newDocument, setNewDocument] = useState({
        title: '',
        description: '',
        documentType: 'notes',
        fileName: '',
        fileType: '',
        fileSize: 0,
        fileUrl: '',
        storageKey: '',
        folder: '',
        tags: '', // comma-separated
        isPublic: false,
        subjectId: '',
    });
    const [isLoading, setIsLoading] = useState(false);

    useEffect(() => {
        if (token) {
            fetchDocuments();
            fetchSubjects();
        }
    }, [token]);

    const fetchDocuments = async () => {
        try {
            const res = await fetch("/api/documents", {
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            });
            if (res.ok) {
                const data = await res.json();
                setDocuments(data);
            } else {
                toast.error("Failed to fetch documents.");
            }
        } catch (error) {
            toast.error("An unexpected error occurred while fetching documents.");
            console.error("Fetch documents error:", error);
        }
    };

    const fetchSubjects = async () => {
        try {
            const subjectsRes = await fetch("/api/subjects", { headers: { Authorization: `Bearer ${token}` } });
            if (subjectsRes.ok) setSubjects(await subjectsRes.json());
            else toast.error("Failed to fetch subjects.");

        } catch (error) {
            toast.error("Error fetching subject dependencies.");
            console.error("Fetch dependencies error:", error);
        }
    };

    const handleInputChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        const { name, value, type, checked } = e.target;
        setNewDocument(prev => ({
            ...prev,
            [name]: type === 'checkbox' ? checked : (type === 'number' ? parseInt(value) : value),
        }));
    };

    const handleSelectChange = (name: string, value: string) => {
        setNewDocument(prev => ({ ...prev, [name]: value }));
    };

    const handleSubmitNewDocument = async (e: React.FormEvent) => {
        e.preventDefault();
        setIsLoading(true);

        const payload = {
            ...newDocument,
            subjectId: newDocument.subjectId || null,
            description: newDocument.description || null,
            fileName: newDocument.fileName || null,
            fileType: newDocument.fileType || null,
            fileSize: newDocument.fileSize || null,
            storageKey: newDocument.storageKey || null,
            folder: newDocument.folder || null,
            tags: newDocument.tags ? newDocument.tags.split(',').map(s => s.trim()).filter(Boolean) : [],
        };

        try {
            const res = await fetch("/api/documents", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: `Bearer ${token}`,
                },
                body: JSON.stringify(payload),
            });

            if (res.ok) {
                toast.success("Document created successfully!");
                setNewDocument({ // Reset form
                    title: '',
                    description: '',
                    documentType: 'notes',
                    fileName: '',
                    fileType: '',
                    fileSize: 0,
                    fileUrl: '',
                    storageKey: '',
                    folder: '',
                    tags: '',
                    isPublic: false,
                    subjectId: '',
                });
                fetchDocuments(); // Refresh list
            } else {
                const errorData = await res.json();
                toast.error(errorData.error || "Failed to create document.");
            }
        } catch (error) {
            toast.error("An unexpected error occurred while creating document.");
            console.error("Create document error:", error);
        } finally {
            setIsLoading(false);
        }
    };

    return (
        <div className="space-y-6">
            <Toaster position="top-center" richColors />

            <Card className="bg-gray-800 border-gray-700">
                <CardHeader>
                    <CardTitle className="text-white">Upload New Document</CardTitle>
                </CardHeader>
                <CardContent>
                    <form onSubmit={handleSubmitNewDocument} className="grid grid-cols-1 md:grid-cols-2 gap-4">
                        <Input name="title" placeholder="Document Title" value={newDocument.title} onChange={handleInputChange} required className="bg-gray-700 border-gray-600 text-white col-span-2" />
                        <Textarea name="description" placeholder="Description (Optional)" value={newDocument.description} onChange={handleInputChange} className="bg-gray-700 border-gray-600 text-white col-span-2" />
                        
                        <div>
                            <label className="block text-gray-300 text-sm font-bold mb-2">Document Type</label>
                            <Select name="documentType" value={newDocument.documentType} onValueChange={(val) => handleSelectChange('documentType', val)} required>
                                <SelectTrigger className="w-full bg-gray-700 border-gray-600 text-white">
                                    <SelectValue placeholder="Select type" />
                                </SelectTrigger>
                                <SelectContent className="bg-gray-700 border-gray-600 text-white">
                                    <SelectItem value="notes">Notes</SelectItem>
                                    <SelectItem value="textbook">Textbook</SelectItem>
                                    <SelectItem value="slides">Slides</SelectItem>
                                    <SelectItem value="previous_paper">Previous Paper</SelectItem>
                                    <SelectItem value="fat_paper">FAT Paper</SelectItem>
                                    <SelectItem value="question_bank">Question Bank</SelectItem>
                                    <SelectItem value="formula_sheet">Formula Sheet</SelectItem>
                                    <SelectItem value="cheat_sheet">Cheat Sheet</SelectItem>
                                    <SelectItem value="other">Other</SelectItem>
                                </SelectContent>
                            </Select>
                        </div>
                        <Input name="fileUrl" placeholder="File URL (e.g., Google Drive link)" value={newDocument.fileUrl} onChange={handleInputChange} required className="bg-gray-700 border-gray-600 text-white" />
                        
                        <div>
                            <label className="block text-gray-300 text-sm font-bold mb-2">Subject (Optional)</label>
                            <Select name="subjectId" value={newDocument.subjectId || ''} onValueChange={(val) => handleSelectChange('subjectId', val)}>
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
                        <Input name="folder" placeholder="Folder (Optional)" value={newDocument.folder} onChange={handleInputChange} className="bg-gray-700 border-gray-600 text-white" />
                        
                        <Input name="tags" placeholder="Tags (comma-separated)" value={newDocument.tags} onChange={handleInputChange} className="bg-gray-700 border-gray-600 text-white col-span-2" />

                        <div className="md:col-span-2 flex items-center space-x-2">
                            <Checkbox
                                id="isPublic"
                                name="isPublic"
                                checked={newDocument.isPublic}
                                onCheckedChange={(checked) => setNewDocument(prev => ({ ...prev, isPublic: checked as boolean }))}
                                className="h-4 w-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
                            />
                            <label htmlFor="isPublic" className="text-gray-300">Make Public?</label>
                        </div>
                        
                        <div className="md:col-span-2">
                            <Button type="submit" disabled={isLoading} className="w-full bg-blue-600 hover:bg-blue-700">
                                {isLoading ? "Uploading..." : "Add Document"}
                            </Button>
                        </div>
                    </form>
                </CardContent>
            </Card>

            <Card className="bg-gray-800 border-gray-700">
                <CardHeader>
                    <CardTitle className="text-white">Your Documents</CardTitle>
                </CardHeader>
                <CardContent>
                    {documents.length === 0 ? (
                        <p className="text-gray-400">No documents found. Upload one above!</p>
                    ) : (
                        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                            {documents.map(doc => (
                                <Card key={doc.id} className="bg-gray-700 border-gray-600 p-4">
                                    <h3 className="font-bold text-lg text-white">{doc.title}</h3>
                                    {doc.description && <p className="text-gray-300 text-sm mt-1">{doc.description}</p>}
                                    <p className="text-gray-300 text-sm">Type: {doc.documentType}</p>
                                    {doc.subjectId && <p className="text-gray-300 text-sm">Subject: {subjects.find(sub => sub.id === doc.subjectId)?.name}</p>}
                                    <p className="text-gray-300 text-sm">URL: <a href={doc.fileUrl} target="_blank" rel="noopener noreferrer" className="text-blue-400 hover:underline">{doc.fileUrl}</a></p>
                                    {doc.tags && doc.tags.length > 0 && <p className="text-gray-300 text-sm">Tags: {doc.tags.join(', ')}</p>}
                                </Card>
                            ))}
                        </div>
                    )}
                </CardContent>
            </Card>
        </div>
    );
}
