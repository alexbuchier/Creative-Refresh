-- Create a function to draw bubbles
function drawBubbles(sw, sh)
    local bubbles = {}
    local tempX, tempY
    for i=1, 2000 do
        tempX = math.random(0, sw)
        tempY = math.random(0, sh)
        bubbles[i] = {x = tempX, y = tempY, radius = math.random(4, 8)}
    end
    return bubbles
end

-- Create a function to draw pink clouds
function drawClouds(sw, sh)
    local clouds = {}
    local tempX, tempY
    for i=1, 2000 do
        tempX = math.random(0, sw)
        tempY = math.random(0, sh)
        clouds[i] = {x = tempX, y = tempY, radius = math.random(2, 4)}
    end
    return clouds
end

-- Create a function to draw rain drops
function drawRainDrops(sw, sh)
    local rainDrops = {}
    local tempX, tempY
    for i=1, 2000 do
        tempX = math.random(0, sw)
        tempY = math.random(0, sh)
        rainDrops[i] = {x = tempX, y = tempY, length = math.random(2, 8)}
    end
    return rainDrops
end

-- Create a function to draw bubbles with pink clouds
function drawBubbleClouds(sw, sh)
    local bubbleClouds = {}
    local tempX, tempY
    for i=1, 2000 do
        tempX = math.random(0, sw)
        tempY = math.random(0, sh)
        bubbleClouds[i] = {x = tempX, y = tempY, radius = math.random(4, 8), cloudRadius = math.random(2, 4)}
    end
    return bubbleClouds
end

-- Create a function to draw lightning
function drawLightning(sw, sh)
    local lightning = {}
    local tempX, tempY
    for i=1, 2000 do
        tempX = math.random(0, sw)
        tempY = math.random(0, sh)
        lightning[i] = {x = tempX, y = tempY, length = math.random(2, 8)}
    end
    return lightning
end

-- Create a function to draw rainbows
function drawRainbow(sw, sh)
    local rainbow = {}
    local tempX, tempY
    for i=1, 2000 do
        tempX = math.random(0, sw)
        tempY = math.random(0, sh)
        rainbow[i] = {x = tempX, y = tempY, radius = math.random(2, 4)}
    end
    return rainbow
end

-- Create a function to draw stars
function drawStars(sw, sh)
    local stars = {}
    local tempX, tempY
    for i=1, 2000 do
        tempX = math.random(0, sw)
        tempY = math.random(0, sh)
        stars[i] = {x = tempX, y = tempY, radius = math.random(1, 2)}
    end
    return stars
end

-- Create a function to draw birds
function drawBirds(sw, sh)
    local birds = {}
    local tempX, tempY
    for i=1, 2000 do
        tempX = math.random(0, sw)
        tempY = math.random(0, sh)
        birds[i] = {x = tempX, y = tempY, size = math.random(2, 4)}
    end
    return birds
end

-- Create a function to draw fireflies
function drawFireflies(sw, sh)
    local fireflies = {}
    local tempX, tempY
    for i=1, 2000 do
        tempX = math.random(0, sw)
        tempY = math.random(0, sh)
        fireflies[i] = {x = tempX, y = tempY, radius = math.random(2, 4)}
    end
    return fireflies
end

-- Create a refresh function
function refresh(sw, sh)
    local bubbles = drawBubbles(sw, sh)
    local clouds = drawClouds(sw, sh)
    local rainDrops = drawRainDrops(sw, sh)
    local bubbleClouds = drawBubbleClouds(sw, sh)
    local lightning = drawLightning(sw, sh)
    local rainbow = drawRainbow(sw, sh)
    local stars = drawStars(sw, sh)
    local birds = drawBirds(sw, sh)
    local fireflies = drawFireflies(sw, sh)
    return bubbles, clouds, rainDrops, bubbleClouds, lightning, rainbow, stars, birds, fireflies
end