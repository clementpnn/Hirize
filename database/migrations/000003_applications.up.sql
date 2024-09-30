CREATE TYPE application_status AS ENUM (
    'pending',
    'followed-up',
    'rejected',
    'interview'
);

CREATE TABLE IF NOT EXISTS applications (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    company_name VARCHAR(255) NOT NULL,
    position VARCHAR(255) NOT NULL,
    application_date DATE NOT NULL,
    status application_status NOT NULL DEFAULT 'pending',
    follow_up_date DATE,
    contact_info VARCHAR(255),
    job_posting_url TEXT,
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_application_user FOREIGN KEY (user_id) REFERENCES users(id)ON DELETE CASCADE
);

CREATE INDEX idx_application_user_id ON applications(user_id);