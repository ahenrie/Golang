# Write your MySQL query statement below
SELECT user_id, concat(UPPER(LEFT(name, 1)),LOWER(SUBSTRING(name, 2))) AS name
FROM Users
ORDER BY 1
