#!/bin/bash
tmux new-session -d -s test -c ~/Documents/GitHub/tmuxcraft
tmux rename-window -t test: nvim
sleep 0.1
tmux send-keys -t test:nvim.1 "nvim" C-m
tmux split-window -t test:nvim.1 -c ~/Documents/GitHub/tmuxcraft -h -l 40
sleep 0.1
tmux send-keys -t test:nvim.2 "tty-clock -t" C-m
tmux split-window -t test:nvim.2 -c ~/Documents/GitHub/tmuxcraft -v -l 20
sleep 0.1
tmux send-keys -t test:nvim.3 "bash" C-m
tmux new-window -t test -c ~/Documents/GitHub/tmuxcraft -n resource
sleep 0.1
tmux send-keys -t test:resource.1 "htop" C-m
tmux new-window -t test -c ~/Documents/GitHub/tmuxcraft -n test
