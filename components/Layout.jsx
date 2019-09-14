import React from "react";
import "../static/styles.css";

const Layout = props => {
  return (
    <div className="main-box">
      <div className="p-4">{props.children}</div>
    </div>
  );
};

export default Layout;
