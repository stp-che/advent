# frozen_string_literal: true

path = ARGV.first

WINNER = {
  rock: :paper,
  scissors: :rock,
  paper: :scissors
}.freeze

SCORES_FOR_CHOICE = { rock: 1, paper: 2, scissors: 3 }.freeze

CHOICE_CODES = {
  'A' => :rock,
  'B' => :paper,
  'C' => :scissors,
  'X' => :rock,
  'Y' => :paper,
  'Z' => :scissors
}.freeze

def my_score(op_choice, my_choice)
  outcome(op_choice, my_choice) + SCORES_FOR_CHOICE[my_choice]
end

def outcome(op_choice, my_choice)
  return 6 if WINNER[op_choice] == my_choice

  return 3 if op_choice == my_choice

  0
end

rounds = File.read(path).split("\n").map(&:split)

scores = rounds.reduce(0) do |sum, round|
  op_choice = CHOICE_CODES[round[0]]
  my_choice = CHOICE_CODES[round[1]]
  sum + my_score(op_choice, my_choice)
end

puts scores
