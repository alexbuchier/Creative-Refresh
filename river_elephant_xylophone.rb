############################################
# Creative Refresh
# By: [Insert Name Here]
############################################

#Declare a variable
refreshing_time = 15

#Create the loop
loop do 
  #Output a prompt
  puts "It's time to get creative. Get up and take a #{refreshing_time} minute break!"
  
  #Decrement the variable
  refreshing_time -= 1
  
  #Break out of the loop
  break if refreshing_time == 0
end

#Output a message
puts "You've done it! Now get back to creating!"

#Create an array
creative_genres = ["writing", "painting", "design", "photography", "music"]

#Iterate through the array with .each
creative_genres.each do |genre|
  puts "The #{genre} genre is perfect for creative refreshment."
end

#Create a hash
creative_tools = { pen: "pencil or pen", paper: "plain or sketchbook", paint: "acrylic, oil, or watercolor" }

#Iterate through hash using .each
creative_tools.each do |tool, type|
  puts "A great tool for creative refreshment is a #{type} #{tool}."
end

#Create a method
def creative_activity
  #Get a random element from the creative_genres array and store it in a variable
  random_genre = creative_genres.sample
  
  puts "Try a creative activity from the #{random_genre} genre."
end

#Call the creative_activity method
creative_activity

#Declare a variable
creative_challenge = "What is one small project you can complete in the next five minutes?"

#Output a message
puts "Here is your creative challenge: #{creative_challenge}"