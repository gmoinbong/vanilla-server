#For start paste in bash - ./start_server.sh

LOCK_FILE="/home/vladyslav/Documents/web-projects/vanilla-server"

# Handle signal SIGINT (Ctrl+C)

cleanup(){
    echo "Cleaning up.."

    if [ -f "$LOCK_FILE" ]; then
        rm "$LOCK_FILE"
        echo "Lock file removed"
    fi
    exit 1
}

trap 'cleanup' SIGINT

# Env CONFIG_PATH

export CONFIG_PATH=/home/vladyslav/Documents/web-projects/vanilla-server/config/local.yaml 

# Go to directory with Docker Compose

cd /home/vladyslav/Documents/web-projects/vanilla-server/docker #Paste docker directory or create folder in root of project 

# Run Docker Compose 

docker-compose up -d


# Run Golang app

cd /home/vladyslav/Documents/web-projects/vanilla-server
go run ./cmd/main.go &

wait