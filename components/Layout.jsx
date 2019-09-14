import React from "react";

const Layout = props => {
  return (
    <div style={{ height: "100vh", backgroundColor: "#f8f8f8" }}>
      <div className="m-4">{props.children}</div>
    </div>
  );
};

export default Layout;
