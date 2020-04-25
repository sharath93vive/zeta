# Exeception command

    Description:
        Remove the duplicate or repetitive continuous words from the given paragraph

    Command : 
        ./fix_repetitions.sh inputs/sample_input.txt

    Implementation:
    
    . using awk, compare each column/word in the current context(line) with the previous column
      If the two-columns does not match, then printing the word in the column.
      The same procedure will be followed until it reaches the end of the line.

      NF = number of field
      OFS = output field separator


