import React from "react";
import { Table } from "react-bootstrap";

const DisplayColumn = props => {
  return (
    <div className="border rounded p-2 shadow display-col">
      <h2 className="title-font">{props.title}</h2>
      <Table striped hover>
        <thead>
          <tr>
            {props.headers.map((item, idx) => {
              return <th key={idx}>{item}</th>;
            })}
          </tr>
        </thead>
        <tbody>
          {props.items.map(item => {
            return (
              <tr key={item.id}>
                {Object.keys(item).map((key, idx) => {
                  if (key !== "id") {
                    return <td key={idx}>{item[key]}</td>;
                  }
                })}
              </tr>
            );
          })}
        </tbody>
      </Table>
    </div>
  );
};

export default DisplayColumn;
