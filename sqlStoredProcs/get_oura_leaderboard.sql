
CREATE DEFINER="doadmin"@"%" PROCEDURE "get_oura_leaderboard"()
BEGIN
   -- Step 1: Gather unique entries
WITH unique_entries AS (
    SELECT 
        jsondata,
        wallet,
        MIN(wallet) AS first_wallet -- Select the first wallet that has this jsondata
    FROM 
        walletchat.ouradata
    WHERE 
        jsondata NOT LIKE '%{"data":[],"next_token":null}%'
        AND jsondata LIKE '%"score":%' -- Ensure the score pattern exists
        AND endpoint = 'daily_sleep' -- Filter for daily_sleep endpoint
    GROUP BY 
        jsondata, wallet -- Group by jsondata and wallet for uniqueness per wallet
),

-- Step 2: Extract scores and rank them by wallet
extracted_scores AS (
    SELECT 
        a.wallet,
        CAST(
            SUBSTRING(
                a.jsondata, 
                LOCATE('"score":', a.jsondata) + 8, 
                LOCATE(',', a.jsondata, LOCATE('"score":', a.jsondata) + 8) - LOCATE('"score":', a.jsondata) - 8
            ) AS UNSIGNED
        ) AS score, -- Extract score dynamically
        ROW_NUMBER() OVER (PARTITION BY a.wallet ORDER BY a.id DESC) AS rn -- Rank by most recent entries per wallet
    FROM 
        unique_entries u
    JOIN 
        walletchat.ouradata a ON u.jsondata = a.jsondata AND u.first_wallet = a.wallet
    WHERE 
        a.endpoint = 'daily_sleep' -- Ensure only daily_sleep entries are included
),

-- Step 3: Group scores for average calculation
avg_sleep_data AS (
    SELECT 
        wallet,
        AVG(score) AS avg_sleep -- Calculate average sleep score
    FROM 
        extracted_scores
    WHERE 
        rn <= 7 -- Take only the first 7 entries per wallet
    GROUP BY 
        wallet
)

-- Step 4: Print the leaderboard
SELECT 
    a.wallet,
    COUNT(u.jsondata) * 5 AS total_points,
    COALESCE(avg_sleep_data.avg_sleep, 0) AS avg_sleep,
    n.name -- Add the name from the addrnameitems table
FROM 
    unique_entries u
JOIN 
    walletchat.ouradata a ON u.jsondata = a.jsondata AND u.first_wallet = a.wallet
LEFT JOIN 
    avg_sleep_data ON a.wallet = avg_sleep_data.wallet
LEFT JOIN 
    walletchat.addrnameitems n ON a.wallet = n.address -- Join addrnameitems for name
WHERE 
    a.endpoint = 'daily_sleep' -- Filter for daily_sleep entries
GROUP BY 
    a.wallet, avg_sleep_data.avg_sleep, n.name
ORDER BY 
    total_points DESC;

END
