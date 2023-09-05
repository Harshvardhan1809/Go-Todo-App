# Go-Todo-App

Learning Golang while creating a presentable project 
Technologies used:
Backend -> Golang
DB -> MySQL
Frontend -> React (TS)

## Important 
For running frontend : npm run dev </br>
</br>
For running backend with logging and live reloading : </br>
alias air="$(go env GOPATH)"/bin/air </br>
air </br>
(run the air command in the directory where main.go lies i.e. src)

## Development process
<ul>
<h6> Cookies and tokens </h6> 
<li>Planned on using cookies for authentication but not working well. Getting the set-cookie header while returning response from request, but the cookie is not displayed in the cookies area of dev-tools </br>In case of Firefox, no set-cookie header at all </li>
<li>Hence, might shift to JWT tokens </li>
