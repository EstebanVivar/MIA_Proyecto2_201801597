import React, { Component } from "react";
import axios from 'axios';

export default class Profile extends Component {

    render() {
        return (

            <div className="container">
            <div className="row justify-content-center">
                <div className="col-12 col-lg-10  mx-auto">
                    <div className="my-4">                       
                        <form>
                            <div className="row mt-5 align-items-center">
                                <div className="col-md-6 text-center mb-5">
                                    <div className="avatar avatar-xl">
                                        <img src="https://bootdey.com/img/Content/avatar/avatar6.png" alt="..." className="avatar-img rounded-circle" />
                                    </div>
                                </div>
                                <div className="col">
                                    <div className="row align-items-center">
                                        <div className="col-md-7">
                                            <h4 className="mb-1">Brown, Asher</h4>
                                            <p className="small mb-3"><span className="badge badge-dark">New York, USA</span></p>
                                        </div>
                                    </div>
                                    <div className="row mb-4">                                        
                                        <div className="col">
                                            <p className="small mb-0 text-muted">Nec Urna Suscipit Ltd</p>
                                            <p className="small mb-0 text-muted">P.O. Box 464, 5975 Eget Avenue</p>
                                            <p className="small mb-0 text-muted">(537) 315-1481</p>
                                        </div>
                                    </div>
                                </div>
                            </div>
                            <hr className="my-4" />
                            <div className="form-row">
                                <div className="form-group col-md-6">
                                    <label for="firstname">Firstname</label>
                                    <input type="text" id="firstname" className="form-control" placeholder="Brown" />
                                </div>
                                <div className="form-group col-md-6">
                                    <label for="lastname">Lastname</label>
                                    <input type="text" id="lastname" className="form-control" placeholder="Asher" />
                                </div>
                            </div>
                            <div className="form-group">
                                <label for="inputEmail4">Email</label>
                                <input type="email" className="form-control" id="inputEmail4" placeholder="brown@asher.me" />
                            </div>
                            <div className="form-group">
                                <label for="inputAddress5">Address</label>
                                <input type="text" className="form-control" id="inputAddress5" placeholder="P.O. Box 464, 5975 Eget Avenue" />
                            </div>
                            
                            <hr className="my-4" />
                            <div className="row mb-4">
                                <div className="col-md-6">
                                    <div className="form-group">
                                        <label for="inputPassword4">Old Password</label>
                                        <input type="password" className="form-control" id="inputPassword5" />
                                    </div>
                                    <div className="form-group">
                                        <label for="inputPassword5">New Password</label>
                                        <input type="password" className="form-control" id="inputPassword5" />
                                    </div>
                                    <div className="form-group">
                                        <label for="inputPassword6">Confirm Password</label>
                                        <input type="password" className="form-control" id="inputPassword6" />
                                    </div>
                                </div>
                                <div className="col-md-6">
                                    <p className="mb-2">Password requirements</p>
                                    <p className="small text-muted mb-2">To create a new password, you have to meet all of the following requirements:</p>
                                    <ul className="small text-muted pl-4 mb-0">
                                        <li>Minimum 8 character</li>
                                        <li>At least one special character</li>
                                        <li>At least one number</li>
                                        <li>Can’t be the same as a previous password</li>
                                    </ul>
                                </div>
                            </div>
                            <button type="submit" className="btn btn-primary">Save Change</button>
                        </form>
                    </div>
                </div>
            </div>
            
            </div>
        )
    }
}

