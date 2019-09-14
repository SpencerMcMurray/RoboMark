import React from "react";
import "../static/styles.css";

const Layout = props => {
  return <div className="main-box">{props.children}</div>;
};

export default Layout;
