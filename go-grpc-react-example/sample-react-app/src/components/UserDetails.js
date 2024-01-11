import React, { Component } from 'react';
import { UsrClient } from '../proto/userInfo_grpc_web_pb';
import { UserRequest } from '../proto/userInfo_pb';// Import the generated client code

class UserDetail extends Component {
    constructor(props) {
        super(props);
        this.state = {
            user: null,
            error: null,
        };
    }

    componentDidMount() {
        const client = new UsrClient('http://localhost:8083'); // Replace with your gRPC server URL

        // Create a request with the user's name
        const request = new UserRequest();
        // Replace with the desired user's name

        // Call the GetUser RPC
        client.getUser(request, {}, (err, response) => {
            if (!err) {
                this.setState({ user: response.getUser() });
            } else {
                this.setState({ error: 'Error fetching user data' });
            }
        });
    }

    render() {
        const { user, error } = this.state;

        return (
            <div>
                {error ? (
                    <p>Error: {error}</p>
                ) : user ? (
                    <div>
                        <h1>User Details</h1>
                        <p>Name: {user.getName()}</p>
                        <p>Age: {user.getAge()}</p>
                        <h2>Address</h2>
                        <p>Street: {user.getAddress().getStreet()}</p>
                        <p>City: {user.getAddress().getCity()}</p>
                        <p>State: {user.getAddress().getState()}</p>
                        <p>Zip: {user.getAddress().getZip()}</p>
                        {/* You can display phone details similarly */}
                        <p>Updated At: {user.getUpdatedAt()}</p>
                        <p>Created At: {user.getCreatedAt()}</p>
                    </div>
                ) : (
                    <p>Loading...</p>
                )}
            </div>
        );
    }
}

export default UserDetail;
