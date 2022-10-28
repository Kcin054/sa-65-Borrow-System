import React, { useState, useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import { List_dataInterface } from "../models/IList_data";
import { GetList_data } from "../Services/HttpClientService";


function List_datas() {
  const [list_datas, setList_datas] = useState<List_dataInterface[]>([]);

  useEffect(() => {
    List_datas();
  }, []);

  const List_datas = async () => {
    let res = await GetList_data();
    if (res) {  
      setList_datas(res);
    } 
  };

  const columns: GridColDef[] = [
    { field: "ID", headerName: "ลำดับ", width: 50 },


   {
      field: "Equipment",
      headerName: "อุปกรณ์",
      width: 200,
      valueFormatter: (params) => params.value.Equipment_name,
    },
    {
      field: "Room",
      headerName: "ห้อง",
      width: 50,
      valueFormatter: (params) => params.value.ID,
    },
    {
      field: "Borrow_time",
      headerName: "วันที่ยืม",
      width: 110,
    },
    {
      field: "Return_time",
      headerName: "วันที่คืน",
      width: 110,
    },
    {
      field: "Borrow_card",
      headerName: "บัตรยืม",
      width: 70,
      valueFormatter: (params) => params.value.ID,
    },

  ];

  return (
    <div>
      <Container maxWidth="md">
        <Box
          display="flex"
          sx={{
            marginTop: 2,
          }}
        >
          <Box flexGrow={1}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              ข้อมูลห้อง
            </Typography>
          </Box>
          <Box>
            <Button
              component={RouterLink}
              to="/List_data/create"
              variant="contained"
              color="primary"
            >
              กลับ
            </Button>
          </Box>
        </Box>
        <div style={{ height: 400, width: "100%", marginTop: "20px" }}>
          <DataGrid
            rows={list_datas}
            getRowId={(row) => row.ID}
            columns={columns}
            pageSize={5}
            rowsPerPageOptions={[5]}
          />
        </div>
      </Container>
    </div>
  );
}

export default List_datas;