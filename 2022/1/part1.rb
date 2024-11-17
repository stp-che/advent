path = ARGV.first

calories = File.read(path).split("\n\n").map do |batch|
  batch.split.map(&:to_i).sum
end

puts calories.max
