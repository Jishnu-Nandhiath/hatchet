import { useState } from 'react';
import { cn } from '@/next/lib/utils';
import { Button } from '@/next/components/ui/button';
import { CheckIcon, Copy } from 'lucide-react';
import CodeStyleRender from './code-render';
import { Link } from 'react-router-dom';
interface CodeBlockProps {
  title?: string;
  language: string;
  value: string;
  className?: string;
  noHeader?: boolean;
  showLineNumbers?: boolean;
  highlightLines?: number[];
  highlightStrings?: string[];
  link?: string;
}

const defaultHighlightLines: number[] = [];
const defaultHighlightStrings: string[] = [];

export function CodeBlock({
  noHeader = false,
  title,
  language,
  value,
  className,
  highlightLines = defaultHighlightLines,
  highlightStrings = defaultHighlightStrings,
  link,
  ...props
}: CodeBlockProps) {
  const [copied, setCopied] = useState(false);

  const copyToClipboard = async () => {
    await navigator.clipboard.writeText(value);
    setCopied(true);
    setTimeout(() => setCopied(false), 2000);
  };

  return (
    <div
      className={cn(
        'relative rounded-md overflow-hidden border border-muted',
        className,
      )}
    >
      {!noHeader && (
        <div className="flex items-center justify-between px-2 bg-muted/50 border-b rounded-t-md">
          <div className="text-xs text-muted-foreground font-mono">
            {link ? (
              <Link to={link} target="_blank" rel="noopener noreferrer">
                {title || language}
              </Link>
            ) : (
              title || language
            )}
          </div>
          <Button
            variant="ghost"
            size="sm"
            onClick={copyToClipboard}
            className="h-8 px-2"
          >
            {copied ? (
              <CheckIcon className="h-4 w-4" />
            ) : (
              <Copy className="h-4 w-4" />
            )}
          </Button>
        </div>
      )}
      <div className={cn('relative')}>
        <pre
          className={cn(
            'p-4 overflow-auto text-sm font-mono bg-muted/30 rounded-b-md max-h-96 scrollbar-thin scrollbar-track-muted scrollbar-thumb-muted-foreground',
          )}
        >
          <CodeStyleRender
            parsed={value}
            language={language}
            highlightLines={highlightLines}
            highlightStrings={highlightStrings}
            {...props}
          />
        </pre>
      </div>
    </div>
  );
}
