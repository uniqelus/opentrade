INSERT INTO users (id, email, first_name, last_name, state, create_time, update_time)
VALUES ($1, $2, $3, $4, $5, $6, $7)
ON CONFLICT (id) DO UPDATE SET
    email = EXCLUDED.email,
    first_name = EXCLUDED.first_name,
    last_name = EXCLUDED.last_name,
    state = EXCLUDED.state,
    update_time = CURRENT_TIMESTAMP;
