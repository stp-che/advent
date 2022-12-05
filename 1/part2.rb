# frozen_string_literal: true

path = ARGV.first

calories = File.read(path).split("\n\n").map do |batch|
  batch.split.map(&:to_i).sum
end

N = 3
top_n = Array.new(N) { 0 }

calories.each do |val|
  i = 0
  i += 1 while i < N && top_n[i] > val
  next if i == N

  left = []
  left = top_n[0..i-1] if i > 0

  top_n = left + [val] + top_n[i..N-2]
end

puts top_n.sum
