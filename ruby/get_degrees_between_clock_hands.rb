def find_degrees_between_hands(time)
  #Match on 01:00, 1:00, 12:00 or 12:00PM and split by the colon
  t = time.match(/\d+:\d{2}/).to_s.split(":")
  #set variables for later use and convert from string to float
  hour = t[0].to_f
  min = t[1].to_f
 
  #Error checking to see if invalid time was entered and abort execution if true
  if hour > 12 || min > 59 then
    abort("Invalid time")
  end
 
  #Get the degrees of the hour hand
  hourDeg = (60*hour+min)/2
  #Get the degrees of the minute hand
  minDeg = 6*min
  #Get the angle of the hour hand in relation to the minute hand
  fullDeg = hourDeg - minDeg
  #Check to see if angle from previous equation is > 180, if so, subtract that value from 360 to get smallest angle
  fullDeg = fullDeg > 180 ? fullDeg -= 360: fullDeg
  #Return formatted string listing input time and angle
  return "The angle of #{time} is: #{fullDeg.abs}"
end
 
#As you can see from the examples below, you can enter the time in various formats.  The last example shows what happens when you enter an invalid time.
a = find_degrees_between_hands("05:20")
puts a
b = find_degrees_between_hands("12:00PM")
puts b
c = find_degrees_between_hands("3:00 PM")
puts c
d = find_degrees_between_hands("6:00PM")
puts d
e = find_degrees_between_hands("9:00P")
puts e
f = find_degrees_between_hands("19:00PM")
puts f

Output:

The angle of 05:20 is: 40.0 
The angle of 12:00PM is: 0.0 
The angle of 3:00 PM is: 90.0 
The angle of 6:00PM is: 180.0 
The angle of 9:00P is: 90.0 
Invalid time 
