# Cooked - A Generate Software Orientation Workshop
<!-- markdownlint-disable MD013 -->

![GitHub Actions CI Status](https://img.shields.io/github/actions/workflow/status/GenerateNU/cooked/ci.yaml?branch=main&logo=github&label=CI)
![License](https://img.shields.io/github/license/GenerateNU/cooked?label=License)

Welcome to the Fall 24 Software Orientation - a hackathon inspired workshop that will give you the opportunity to walk the basic stack of a typical Generate project.

## Goal

Teams of 8, over a 1.5 hour period, will collaborate to implement the backend and frontend components of a recipe application.
This repository, as is, provides the necessary foundation for all critical features that comprise the final deliverables.
Each team, at the end of the workshop, will be alloted 2 minutes to present their final product to a panel of judges who will grade based on the requirements and additional criteria listed below.

Requirements:

1. You must build two endpoints: **GET** (fetch) all recipes and **POST** (create) a recipe
2. You must implement a basic frontend (inspired, but not bound, by this [Figma](https://www.figma.com/design/FUc9oy4M56lHLYMGX7kOyZ/Engineering-Orientation?node-id=0-1&t=uy34xrVmav24AfgJ-1)) and pull data from the two aformentioned endpoints
3. Everyone must contribute (dock points if we notice people are sitting to the side)

Additional Criteria:

1. Humor - we all need a good laugh to get us through the first week of school
2. Additional Features - did your team create additional endpoints, enhance the frontend, etc.

Final Note:

- You will find a series of comments prefixed by **TODO** scattered across the repository. These indicate small tasks + where code should be written to meet the requirements.

## Tech Stack

- The backend is written in **Go** and uses [**Fiber**](https://docs.gofiber.io/) as a web framework
- The frontend used **Next.js** (React framework) and **Tanstack Query** (formerly React Query) for communicating with the backend
- The database is a **Postgres** instance hosted on Supabase and uses [**sqlx**](https://github.com/jmoiron/sqlx) for transactions

## Usage

1. [Install Nix](https://zero-to-nix.com/start/install)
2. Navigate to the repository in your terminal (use the `cd` command!)
3. Activate the development environment by running the following:

      ```sh
      nix develop --impure
      ```

   You will now have all of the tools you need to start building, including **Go**, **Node.js**, and **pnpm**!

4. Notice the list of commands available within the development environment. These are designed to simplify the development process.
   - Run the following to start the backend server:

        ```sh
        backend-dev
        ```

   - Run the following to start the frontend server:

        ```sh
        frontend-dev
        ```

   - Run the following to display a list of available commands:

        ```sh
        env-help
        ```
