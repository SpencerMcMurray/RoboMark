import React from "react";
import "../static/styles.css";

const Layout = props => {
  return (
    <div
      style={{
        height: "100vh",
        backgroundColor: "#f8f8f8",
        overflowY: "overlay"
      }}
    >
      <div className="p-4">{props.children}</div>
    </div>
  );
};

export default Layout;
