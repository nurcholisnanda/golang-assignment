-- Connect to the mezink database
\c mezink;

-- Insert mock data into the students table
INSERT INTO students (name, marks, created_at) VALUES
    ('John Doe', '{90, 85, 95}', NOW() - INTERVAL '5 days'),
    ('Jane Smith', '{80, 75, 88}', NOW() - INTERVAL '10 days'),
    ('Alice Johnson', '{95, 88, 92}', NOW() - INTERVAL '15 days'),
    ('Bob Anderson', '{85, 90, 87}', NOW() - INTERVAL '20 days'),
    ('Emma Brown', '{92, 89, 91}', NOW() - INTERVAL '25 days'),
    ('Michael White', '{88, 82, 90}', NOW() - INTERVAL '30 days'),
    ('Sophia Clark', '{91, 85, 89}', NOW() - INTERVAL '35 days'),
    ('William Taylor', '{87, 92, 90}', NOW() - INTERVAL '40 days'),
    ('Olivia Wilson', '{89, 86, 93}', NOW() - INTERVAL '45 days'),
    ('James Harris', '{86, 91, 88}', NOW() - INTERVAL '50 days');
