# AccelDatabase

## Building Database

1. Update/Create local .env file. Use the .env.example as a template.

   - This file will contain your github personal access token. The personal access token will need access to the AccelDatabase repo.
2. Run Instead git command to enable go lang to access private repos

   - Run local comand
     `git config --global url."https://${USERNAME}:${PERSONAL_ACCESS_TOKEN}@github.com/".insteadOf "https://github.com/"`
   - Replace USERNAME with your GitHub username and PERSONAL_ACCESS_TOKEN with your actual personal access token. This command tells Git to use your personal access token when accessing GitHub repositories instead of the default HTTPS URL, which is necessary for pulling from private repositories without interactive authentication
3. Build the postgress database image

   - Run dockerfile (This will build the postgress database image):

   `docker build -t accel_db_img .`

   - The dockerfile contains database username and password
4. Create and Run database container

   - Run docker-compose.yml
     `docker-compose -f docker-compose.yml up`
   - The yml file contiains
     - Hostname
     - Port
     - Database name
     - Database username
     - Database password


## Run Go Lang API service

1