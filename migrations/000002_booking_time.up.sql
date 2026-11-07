ALTER TABLE core.booking ADD CONSTRAINT booking_time_order CHECK (start_time < end_time);
