-- Seed para popular a tabela transactions com algumas transações de exemplo
INSERT INTO transactions (id, account_id_from, account_id_to, amount, created_at) VALUES
  ('770e8400-e29b-41d4-a716-446655440001', '660e8400-e29b-41d4-a716-446655440001', '660e8400-e29b-41d4-a716-446655440002', 25.00, NOW()),
  ('770e8400-e29b-41d4-a716-446655440002', '660e8400-e29b-41d4-a716-446655440002', '660e8400-e29b-41d4-a716-446655440003', 15.50, NOW()),
  ('770e8400-e29b-41d4-a716-446655440003', '660e8400-e29b-41d4-a716-446655440003', '660e8400-e29b-41d4-a716-446655440001', 30.00, NOW()); 