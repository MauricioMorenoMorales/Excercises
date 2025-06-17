

SELECT `name`, `population`, `area` FROM `World` WHERE area >= 3000000 OR population >= 25000000;

-- Low fat definition: low_fats = 'Y'
-- recyclable definition: reyclable = 'Y'

SELECT `product_id` from `Products` WHERE low_fats = 'Y' AND recyclable = 'Y';

SELECT `name` from `Customer` WHERE referee_id IS NULL OR referee_id != 2;

SELECT
  employee_id,
  IF(employee_id % 2 = 1 AND name NOT REGEXP '^M', salary, 0) AS bonus
FROM
  employees
ORDER BY
  employee_id

-- Otra forma de escribir esto mismo
SELECT employee_id,
    CASE WHEN
        employee_id % 2 <> 0 AND name NOT LIKE 'M%' THEN salary
        ELSE 0
    END AS bonus
FROM Employees
ORDER BY employee_id