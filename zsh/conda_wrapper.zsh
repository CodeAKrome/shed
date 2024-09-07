# If ai_commit.py is on your path and you want to prefix the rest of your arguments.
# This assumes you named the conda environment the same base name as the python script.
# Example:
# alias aicom='noglob condame "ai_commit" "--service ollama"'

# Wrap command in conda environment call and allow an optional prefix before the rest of the arguments.
condame() {
    local base_name="$1"
    local script_name="${base_name}.py"
    shift  # Remove the first argument (base name) from the argument list
    conda activate "$base_name" || return 1

    # Check for optional prefix argument
    if [[ -n "$2" ]]; then
        prefix="$2"
        shift  # Remove its value from the argument list
	"$script_name" "$prefix" "$@"
    else
	"$script_name" "$@"
    fi

    # Capture the exit status of the Python script
    local exit_status=$?
    conda deactivate
    # Return the exit status of the Python script
    return $exit_status
}
