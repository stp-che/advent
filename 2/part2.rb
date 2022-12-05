# frozen_string_literal: true

path = ARGV.first

WINNER = {
  rock: :paper,
  scissors: :rock,
  paper: :scissors
}.freeze

LOSER = WINNER.invert.freeze

SCORES_FOR_CHOICE = { rock: 1, paper: 2, scissors: 3 }.freeze

CHOICE_CODES = {
  'A' => :rock,
  'B' => :paper,
  'C' => :scissors
}.freeze

OUTCOME_CODES = { 'X' => 0, 'Y' => 3, 'Z' => 6 }.freeze

def my_score(op_choice, expected_outcome)
  my_choice = get_my_choice(op_choice, expected_outcome)
  expected_outcome + SCORES_FOR_CHOICE[my_choice]
end

def get_my_choice(op_choice, expected_outcome)
  case expected_outcome
  when 0
    LOSER[op_choice]
  when 3
    op_choice
  when 6
    WINNER[op_choice]
  end
end

rounds = File.read(path).split("\n").map(&:split)

scores = rounds.reduce(0) do |sum, round|
  op_choice = CHOICE_CODES[round[0]]
  expected_outcome = OUTCOME_CODES[round[1]]
  sum + my_score(op_choice, expected_outcome)
end

puts scores
