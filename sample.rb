# frozen_string_literal: true

# this is a sample piece of Ruby code to test out ktags

class Foo
  def initialize(arg1, arg2)
    puts arg1
    puts arg2
    puts 'this is the initializer'
  end

  def call
    %w[1 2 3].each do |num_str|
      puts num_str
    end
  end
end
