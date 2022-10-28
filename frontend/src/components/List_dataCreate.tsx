import React, { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import Button from "@mui/material/Button";
import FormControl from "@mui/material/FormControl";
import Container from "@mui/material/Container";
import Paper from "@mui/material/Paper";
import Grid from "@mui/material/Grid";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import Divider from "@mui/material/Divider";
import Snackbar from "@mui/material/Snackbar";
import Select, { SelectChangeEvent } from "@mui/material/Select";
import MuiAlert, { AlertProps } from "@mui/material/Alert";
import TextField from "@mui/material/TextField";
import { AdapterDateFns } from "@mui/x-date-pickers/AdapterDateFns";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { DatePicker } from "@mui/x-date-pickers/DatePicker";
import SaveIcon from '@mui/icons-material/Save';
import { EquipmentInterface } from "../models/IEquipment";
import { Borrow_cardInterface } from "../models/IBorrow_card";
import { List_dataInterface } from "../models/IList_data";
import { RoomInterface } from "../models/IRoom";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import {
  GetEquipment,
  GetBorrow_card,
  GetRoom,
  CreateList_datas,
  GetList_data,
  
} from "../Services/HttpClientService";
import FormHelperText from "@mui/material/FormHelperText";
import { FurnitureInterface } from "../models/IFurniture";

const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
  props,
  ref
) {
  return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

function List_dataCreate() {

  // const columns: GridColDef[] = [
  //   { field: "ID", headerName: "ลำดับ", width: 50 },


  //   {
  //     field: "Furniture_type",
  //     headerName: "เฟอนิเจอร์",
  //     width: 150,
  //     valueFormatter: (params) => params.value.ID,
  //   },
  //   {
  //     field: "Set_of_furniture_id",
  //     headerName: "เซทเฟอนิเจอร์",
  //     width: 150,
  //     valueFormatter: (params) => params.value.Set_of_furniture_title,
  //   },
  // ];
  


   //const [furnitures, setFurnitures] = useState<FurnitureInterface[]>([]);//
  const [Equipment, setEquipment] = useState<EquipmentInterface[]>([]);
  const [Borrow_card, setBorrow_card] = useState<Borrow_cardInterface[]>([]);
  const [Room, setRoom] = useState<RoomInterface[]>([]);
  const [list_data, setList_data] = useState<List_dataInterface>({
    Borrow_time: new Date(),
    Return_time: new Date(),
  });
  
 

  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(false);

  const handleClose = (
    event?: React.SyntheticEvent | Event,
    reason?: string
  ) => {
    if (reason === "clickaway") {
      return;
    }
    setSuccess(false);
    setError(false);
  };

  const handleChange = (event: SelectChangeEvent) => {
    const name = event.target.name as keyof typeof List_dataCreate;
    setList_data({
      ...list_data,
      [name]: event.target.value,
    });

    
  };


  

  const getEquipment = async () => {
    let res = await GetEquipment();
    if (res) {
      setEquipment(res);
    }
  };

  const getBorrow_card = async () => {
    let res = await GetBorrow_card();
    if (res) {
      setBorrow_card(res);
    }
  };

  const getRoom = async () => {
    let res = await GetRoom();
    if (res) {
      setRoom(res);
    }
  };
  // const Furnitures = async () => {
  //   let res = await GetFurniture();
  //   if (res) {
  //     setFurnitures(res);
  //   } 
  // };

  //



  useEffect(() => {
    getEquipment();
    getBorrow_card();
    getRoom();
    //Furnitures();
   
  }, []);

  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  async function submit() {
    let data = {
      Equipment_id: convertType(list_data.Equipment_id),
      Borrow_card_id: convertType(list_data.Borrow_card_id),
      Room_id: convertType(list_data.Room_id),
      Borrow_time: list_data.Borrow_time,
      Return_time: list_data.Return_time,


    };

    //
    console.log(data);
    let res = await CreateList_datas(data);
   
    // console.log(res);
    if (res) { 
      setSuccess(true);
    } else {
      setError(true);
    }

  }

  return (
    <Container maxWidth="md"
    >
      <Snackbar
        open={success}
        autoHideDuration={3000}
        onClose={handleClose}
        anchorOrigin={{ vertical: "top", horizontal: "center" }}
      >
        <Alert onClose={handleClose} severity="success">
          บันทึกข้อมูลสำเร็จ
        </Alert>
      </Snackbar>
      <Snackbar
        open={error}
        autoHideDuration={6000}
        onClose={handleClose}
        anchorOrigin={{ vertical: "top", horizontal: "center" }}
      >
        <Alert onClose={handleClose} severity="error">
          บันทึกข้อมูลไม่สำเร็จ
        </Alert>
      </Snackbar>
      <Paper
      sx ={{ bgcolor :"#E3E3E3"}}>
        
        <Box
          display="flex"
          sx={{
            marginTop: 2,
          }}
        >
          <Box sx={{ paddingX: 2, paddingY: 1 }}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              บันทึกข้อมูลรายการยืม
            </Typography>
          </Box>
        </Box>
        <Divider />
        <Grid container spacing={3} sx={{ padding: 2 }}>
          <Grid item xs={12}>
            <FormControl fullWidth variant="outlined" >
              <p>ประเภทอุปกรณ์</p>
              <Select
                native
                value={list_data.Equipment_id + ""}
                onChange={handleChange}
                inputProps={{
                  name: "Equipment_id",
                  
                  
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกประเภทอุปกรณ์
                </option>
                {Equipment.map((item: EquipmentInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Equipment_name}
                  </option>
                ))}
              </Select>
              
            </FormControl>
          </Grid>
          <Grid item xs={12}>
            <FormControl fullWidth variant="outlined">
              <p>ห้อง</p>
              <Select
                native
                value={list_data.Room_id + ""}
                onChange={handleChange}
                inputProps={{
                  name: "Room_id",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกห้อง
                </option>
                {Room.map((item: RoomInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.ID}
                  </option>
                ))}
                
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>วันที่ยืม</p>
              <LocalizationProvider dateAdapter={AdapterDateFns}>
                <DatePicker
                  value={list_data.Borrow_time}
                  onChange={(newValue) => {
                    setList_data({
                      ...list_data,
                      Borrow_time: newValue,
                    });
                  }}
                  renderInput={(params) => <TextField {...params} />}
                />
              </LocalizationProvider>
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>วันที่คืน</p>
              <LocalizationProvider dateAdapter={AdapterDateFns}>
                <DatePicker
                  value={list_data.Return_time}
                  onChange={(newValue) => {
                    setList_data({
                      ...list_data,
                      Return_time: newValue,
                    });
                  }}
                  renderInput={(params) => <TextField {...params} />}
                />
              </LocalizationProvider>
            </FormControl>
          </Grid>
          <Grid item xs={12}>
            <FormControl fullWidth variant="outlined" >
              <p>บัตรยืม</p>
              <Select
                native
                value={list_data.Borrow_card_id + ""}
                onChange={handleChange}
                inputProps={{
                  name: "Borrow_card_id",
                  
                  
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกบัตรยืม
                </option>
                {Borrow_card.map((item: Borrow_cardInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.ID}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={12}>
            <Button
              component={RouterLink}
              to="/List_datas"
              variant="contained"
              color="inherit"


            >
              ข้อมูลการยืม
            </Button>
            <Button
              style={{ float: "right" }}
              onClick={submit}
              variant="contained"
              color="primary"
              startIcon={<SaveIcon />}
            >
              บันทึก
            </Button>
          </Grid>
        </Grid>
      </Paper>
      
      {/* <div style={{ height: 400, width: "100%", marginTop: "20px" }}>
          <DataGrid
          
            rows={furnitures }
            getRowId={(row) => row.ID}
            columns={columns}
            pageSize={5}
            rowsPerPageOptions={[5]}
          />
        </div> */}

    </Container>
  );
}

export default List_dataCreate;