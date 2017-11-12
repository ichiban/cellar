DELETE FROM accounts;
INSERT INTO accounts (id, name) VALUES
  (1, 'account 1'),
  (2, 'account 2');

DELETE FROM bottles;
INSERT INTO bottles (id, account_id, name, vineyard, varietal, vintage, color, sweetness, country, region, review, rating) VALUES
  (100, 1, 'Number 8', 'Asti Winery', 'Merlot', 2012, 'red', 1, 'USA', 'CA', 'Great value', 4),
  (101, 1, 'Mourvedre', 'Rideau', 'Mourvedre', 2012, 'red', 1, 'USA', 'CA', 'Good but expensive', 3),
  (102, 1, 'Blue''s Cuvee', 'Longoria', 'Cabernet Franc with Merlot, Malbec, Cabernet Sauvignon and Syrah', 2012, 'red', 1, 'USA', 'CA', 'Favorite', 5),
  (200, 2, 'Blackstone Merlot', 'Blackstone', 'Merlot', 2012, 'red', 1, 'USA', 'CA', 'OK', 3),
  (201, 2, 'Wild Horse', 'Wild Horse', 'Pinot Noir', 2010, 'red', 1, 'USA', 'CA', 'Solid Pinot', 4);
