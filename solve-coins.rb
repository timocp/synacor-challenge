#! /usr/bin/env ruby

coins = [[2, 'red'], [3, 'corroded'], [5, 'shiny'], [7, 'concave'], [9, 'blue']]

coins.permutation.each do |a, b, c, d, e|
  res = a[0] + b[0] * c[0]**2 + d[0]**3 - e[0]
  next unless res == 399
  puts "use #{a[1]} coin"
  puts "use #{b[1]} coin"
  puts "use #{c[1]} coin"
  puts "use #{d[1]} coin"
  puts "use #{e[1]} coin"
end
