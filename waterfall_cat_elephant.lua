--[[
This code is intended to provide a 'Creative Refresh' for a single Lua file. It will add a number of creatively written snippets which may provide a sense of newness or variance to a given codebase.

--Make a function to generate a random number
function randomNum(min, max)
	return math.random(min, max)
end

--Create a function to create an array of random numbers
function randomArray(len)
	local arr = {}
	
	for i=1, len do
		table.insert(arr, randomNum(1, 10))
	end
	
	return arr
end

--Create a function to calculate the maximum number in an array
function maxNumInArray(array)
    local maxNum = array[1]
    for i=2, #array do
		if array[i] > maxNum then
			maxNum = array[i]
		end
    end
	
    return maxNum
end

--Create a function to calculate the sum of an array
function arraySum(array)
	local sum = 0
	
	for i=1, #array do
		sum = sum + array[i]
	end
	
	return sum
end

--Create a function to check if a value is in an array
function contains(array, value)
	for i=1, #array do
		if array[i] == value then
			return true
		end
	end
	
	return false
end

--Create a function to check if two arrays are equal
function arraysEqual(arr1, arr2)
	if #arr1 ~= #arr2 then
		return false
	end
	
	for i=1, #arr1 do
		if arr1[i] ~= arr2[i] then
			return false
		end
	end
	
	return true
end

--Create a function to calculate the average of an array
function arrayAverage(array)
	return sum(array) / #array
end

--Create a function to check if a number is a prime
function isPrime(num)
	if num < 2 then
		return false
	end
	
	for i=2, math.floor(math.sqrt(num)) do
		if num % i == 0 then
			return false
		end
	end
	
	return true
end

--Create a function to randomize the order of an array
function randomizeArray(array)
	local newArray = {}
	local tmpArr = table.copy(array)
	
	while #tmpArr > 0 do
		local randIndex = math.random(1, #tmpArr)
		local randElem = tmpArr[randIndex]
		table.remove(tmpArr, randIndex)
		table.insert(newArray, randElem)
	end
	
	return newArray
end

--Create a function to determine if a given number is even
function isEven(num)
	return num % 2 == 0
end

--Create a function to calculate the factorial of a given number
function factorial(num)
	if num == 0 then
		return 1
	end
	
	local result = num
	for i=2, num - 1 do
		result = result * i
	end
	
	return result
end

--Create a function to determine if a given string is a palindrome
function isPalindrome(str)
	return str == string.reverse(str)
end

--Create a function to find the number of words in a string
function countWords(str)
	local matchingWords = {}
	
	for match in string.gmatch(str, "%S+") do
		table.insert(matchingWords, match)
	end
	
	return #matchingWords
end
]]