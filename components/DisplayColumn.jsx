import React from "react";
import { Table } from "react-bootstrap";
import SelectBtn from "../components/buttons/SelectBtn";
import QuestionBtn from "../components/buttons/QuestionBtn";

const DisplayColumn = props => {
  return (
    <div className="border rounded p-2 shadow display-col .bg-light">
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
                <React.Fragment>
                  <td>{item.name ? item.name : item.question}</td>
                  <td>
                    {item.name ? (
                      <SelectBtn id={item.id} />
                    ) : (
                      <QuestionBtn
                        onClick={() => {
                          props.setCurrentItem(item);
                          props.onShowPreview();
                        }}
                        item={item}
                      />
                    )}
                  </td>
                </React.Fragment>
              </tr>
            );
          })}
        </tbody>
      </Table>
    </div>
  );
};

export default DisplayColumn;
