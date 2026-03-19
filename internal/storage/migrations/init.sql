CREATE TABLE IF NOT EXISTS public.courses (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name character varying(255) NOT NULL,
    description text,
    updated timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    logo_path character varying(512),
    is_published boolean DEFAULT false NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.chapters (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    course_id uuid NOT NULL,
    name character varying(255) NOT NULL,
    description text,
    updated timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "position" integer DEFAULT 0 NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT chapters_course_id_fkey FOREIGN KEY (course_id) REFERENCES public.courses(id) ON DELETE CASCADE
);

CREATE INDEX idx_chapters_course_id ON public.chapters USING btree (course_id);

CREATE TABLE IF NOT EXISTS public.lessons (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    chapter_id uuid,
    title text NOT NULL,
    text text NOT NULL,
    updated timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    CONSTRAINT lessons_chapter_id_fkey FOREIGN KEY (chapter_id) REFERENCES public.chapters(id)
);