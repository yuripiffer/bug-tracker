import { useEffect, useState } from 'react';
import Image from 'next/image';

const loadingGifs = [
    '/loading/loading1.gif',
    '/loading/loading2.gif',
    '/loading/loading3.gif',
    '/loading/loading4.gif',
];

interface LoadingScreenProps {
    onLoadingComplete: () => void;
    testMode?: boolean;
}

export default function LoadingScreen({ onLoadingComplete, testMode = false }: LoadingScreenProps) {
    const [progress, setProgress] = useState(0);
    const [currentGif, setCurrentGif] = useState(loadingGifs[0]);

    useEffect(() => {
        if (testMode) {
            setTimeout(() => {
                onLoadingComplete();
            }, 0);
            return;
        }

        const progressInterval = setInterval(() => {
            setProgress(prev => {
                if (prev >= 100) {
                    clearInterval(progressInterval);
                    setTimeout(() => {
                        onLoadingComplete();
                    }, 0);
                    return 100;
                }
                return prev + 5;
            });
        }, 100);

        const gifInterval = setInterval(() => {
            setCurrentGif(prev => {
                const currentIndex = loadingGifs.indexOf(prev);
                const nextIndex = (currentIndex + 1) % loadingGifs.length;
                return loadingGifs[nextIndex];
            });
        }, 1000);

        return () => {
            clearInterval(progressInterval);
            clearInterval(gifInterval);
        };
    }, [testMode, onLoadingComplete]);

    if (testMode) {
        return null;
    }

    return (
        <div data-testid="loading-screen" className="fixed inset-0 bg-gray-900 flex flex-col items-center justify-center">
            <div className="relative w-32 h-32 mb-8">
                <Image 
                    src={currentGif}
                    alt="Loading..."
                    fill
                    priority
                    sizes="128px"
                />
            </div>
            
            <div className="w-64 bg-gray-700 rounded-full h-4 overflow-hidden">
                <div 
                    className="bg-blue-500 h-full transition-all duration-100 ease-out"
                    style={{ width: `${progress}%` }}
                />
            </div>
            
            <p className="text-white mt-4 font-medium">
                Loading... {progress}%
            </p>
        </div>
    );
} 