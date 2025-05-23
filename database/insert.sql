INSERT INTO "user" (id, name, document, profile, login, password, token) VALUES
  ('3fa85f64-5717-4562-b3fc-2c963f66afa6', 'ACME Corp', '1234567890001', 0, 'acme', 'senha123', 'token-acme'),
  ('4fa85f64-5717-4562-b3fc-2c963f66afa6', 'Save the Trees', '9876543210001', 1, 'trees', 'senha123', 'token-trees'),
  ('5fa85f64-5717-4562-b3fc-2c963f66afa6', 'City Hall', '112233445566', 2, 'cityhall', 'senha123', 'token-cityhall'),
  ('6fa85f64-5717-4562-b3fc-2c963f66afa6', 'John Doe', '556677889900', 3, 'johndoe', 'senha123', 'token-johndoe');

INSERT INTO problem (id, title, description, localization) VALUES
  ('a1a1a1a1-1111-1111-1111-111111111111', 'Potholes on Main Street', 'Multiple large potholes causing accidents', 'Main Street, Downtown'),
  ('b2b2b2b2-2222-2222-2222-222222222222', 'Lack of Public Lighting', 'Area is very dark and unsafe at night', '5th Avenue, East Side');


INSERT INTO solution (id, title, description, estimated_cost, likes, deslikes, problem_id) VALUES
  ('c3c3c3c3-3333-3333-3333-333333333333', 'Asphalt Repair', 'Use durable asphalt mix to fill potholes', 15000.00, 10, 2, 'a1a1a1a1-1111-1111-1111-111111111111'),
  ('d4d4d4d4-4444-4444-4444-444444444444', 'Install LED Street Lights', 'Energy-efficient lighting for safer roads', 22000.00, 25, 1, 'b2b2b2b2-2222-2222-2222-222222222222');


INSERT INTO sector (name) VALUES
  ('Infrastructure'),
  ('Public Safety'),
  ('Urban Planning'),
  ('Environment');

  INSERT INTO problem_sector (problem_id, sector_id) VALUES
  ('a1a1a1a1-1111-1111-1111-111111111111', 1), -- Infrastructure
  ('b2b2b2b2-2222-2222-2222-222222222222', 2); -- Public Safety

