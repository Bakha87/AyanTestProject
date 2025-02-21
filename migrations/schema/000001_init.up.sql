CREATE TABLE events (
    id SERIAL PRIMARY KEY,
    order_type VARCHAR(50) NOT NULL,
    session_id UUID NOT NULL,
    card VARCHAR(20) NOT NULL,
    event_date TIMESTAMP WITH TIME ZONE NOT NULL,
    website_url TEXT NOT NULL
);