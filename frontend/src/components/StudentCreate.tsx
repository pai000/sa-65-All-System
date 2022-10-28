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
//import { AdapterDateFns } from "@mui/x-date-pickers/AdapterDateFns";
//import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
//import { DatePicker } from "@mui/x-date-pickers/DatePicker";

import { RoleInterface } from "../interfaces/IRole"; 
import { ProgramInterface } from "../interfaces/IProgram"; 
import { ProvincesInterface } from "../interfaces/IProvince"; 
import { GendersInterface } from "../interfaces/IGender";

import {
    GetGenders,
    GetProvinces,
    GetRole,
    GetProgram,
    CreateStudent,
  } from "../services/HttpClientService";
import { StudentInterface } from "../interfaces/IStudent";


  const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
    props,
    ref
  ) {
    return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
  });
  
  function StudentCreate() {
    const [role, setRole] = useState<RoleInterface[]>([]);
    const [program, setProgram] = useState<ProgramInterface[]>([]);
    const [gender, setGenders] = useState<GendersInterface[]>([]);
    const [provinces, setProvinces] = React.useState<ProvincesInterface[]>([]);
    const [student, setStudent] = useState<Partial<StudentInterface>>({});
  
    const [success, setSuccess] = useState(false);
    const [error, setError] = useState(false);
  
     //เปิดปิดตัว alert
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
      //combobox
      const handleChange = (event: SelectChangeEvent) => {
      const name = event.target.name as keyof typeof StudentCreate;
      setStudent({
        ...student,
        [name]: event.target.value,
      });
    };
  
    //text field
    const handleInputChange = (
      event: React.ChangeEvent<{ id?: string; value: any }>
    ) => {
      const id = event.target.id as keyof typeof student;
      const { value } = event.target;
      setStudent({ ...student, [id]: value });
    };

    const getRoles = async () => {
        let res = await GetRole();
        if (res) {
          setRole(res);
        }
    };

    const getProgram = async () => {
        let res = await GetProgram();
        if (res) {
          setProgram(res);
        }
    };
    
    const getGenders = async () => {
      let res = await GetGenders();
      if (res) {
        setGenders(res);
      }
    };

    const getProvinces = async () => {
      let res = await GetProvinces();
      if (res) {
        setProvinces(res);
      }
    };
  
    useEffect(() => {
      getGenders();
      getRoles();
      getProgram();
      getProvinces();
    }, []);
  
    const convertType = (data: string | number | undefined) => {
      let val = typeof data === "string" ? parseInt(data) : data;
      return val;
    };
  
  
    
    async function submit() {
      let data = {
        GenderID: convertType(student.GenderID),
        ProvinceID: convertType(student.ProvinceID),
        ProgramID: convertType(student.ProgramID),
        RoleID: convertType(student.RoleID),
        STUDENT_NUMBER: student.STUDENT_NUMBER,
        STUDENT_NAME: student.STUDENT_NAME,
        PERSONAL_ID: student.PERSONAL_ID,
        Password: student.Password,
      };
      console.log(data);
  
      let res = await CreateStudent(data);
      console.log(res);
      if (res) { 
        setSuccess(true);
      } else {
        setError(true);
      }
    }
  
    return (
      <Container maxWidth="md">
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
        <Paper  sx ={{ bgcolor :"#E3E3E3"}}>
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
                สร้างข้อมูลนักศึกษา
              </Typography>
            </Box>
          </Box>
          <Divider />
          <Grid container spacing={3} sx={{ padding: 2 }}>

          <Grid item xs={6}>
              <p>รหัสนักศึกษา</p>
              <FormControl fullWidth variant="outlined">
                <TextField
                  id="STUDENT_NUMBER"
                  variant="outlined"
                  type="string"
                  size="medium"
                  placeholder="กรุณากรอกรหัสนักศึกาา"
                  value={student.STUDENT_NUMBER || ""}
                  onChange={handleInputChange}
                />
              </FormControl>
            </Grid>   
  
          <Grid item xs={6}>
              <p>ชื่อ-สกุล</p>
              <FormControl fullWidth variant="outlined">
                <TextField
                  id="STUDENT_NAME"
                  variant="outlined"
                  type="string"
                  size="medium"
                  placeholder="กรุณากรอกชื่อ-สกุล"
                  value={student.STUDENT_NAME || ""}
                  onChange={handleInputChange}
                />
              </FormControl>
            </Grid>
  
            <Grid item xs={6}>
              <p>บัตรประจำตัวประชาชน</p>
              <FormControl fullWidth variant="outlined">
                <TextField
                  id="PERSONAL_ID"
                  variant="outlined"
                  type="string"
                  size="medium"
                  placeholder="กรุณากรอกบัตรประชาชน"
                  value={student.PERSONAL_ID || ""}
                  onChange={handleInputChange}
                  inputProps={{maxLength :13}}
                />
              </FormControl>
            </Grid>
  
            {/* <Grid item xs={6}>
              <p>เลขบัตรประชาชน</p>
              <FormControl fullWidth variant="outlined">
                <TextField
                  id="Civ"
                  variant="outlined"
                  //inputProps={{ inputMode: 'numeric', pattern: '[0-9]*' }}
                  type="string"
                  size="medium"
                  placeholder="กรุณากรอกเลขบัตรประชาชน"
                  value={user?.Civ || ""}
                  onChange={handleInputChange}
                  onKeyPress={(e) => {
                    if (!/[0-9]/.test(e.key)){
                      e.preventDefault()
                    }
                  }}
                  inputProps={{maxLength :13}}
                />
              </FormControl>
            </Grid> */}
  
  
            <Grid item xs={6}>
              <p>รหัสผ่าน</p>
              <FormControl fullWidth variant="outlined">
                <TextField
                  id="Password"
                  variant="outlined"
                  type="string"
                  size="medium"
                  placeholder="กรุณากรอกรหัสผ่าน"
                  value={student.Password || ""}
                  onChange={handleInputChange}
                />
              </FormControl>
            </Grid>
  
  
            <Grid item xs={6}>
              <FormControl fullWidth variant="outlined">
                <p>เพศ</p>
                <Select
                  native
                  value={student.GenderID + ""}
                  onChange={handleChange}
                  inputProps={{
                    name: "GenderID",
                  }}
                >
                  <option aria-label="None" value="">
                    กรุณาเลือกเพศ
                  </option>
                  {gender.map((item: GendersInterface) => (
                    <option value={item.ID} key={item.ID}>
                      {item.GENDER_NAME}
                    </option>
                  ))}
                </Select>
              </FormControl>
            </Grid>
  
  
            <Grid item xs={6}>
              <FormControl fullWidth variant="outlined">
                <p>จังหวัด</p>
                <Select
                  native
                  value={student.ProvinceID + ""}
                  onChange={handleChange}
                  inputProps={{
                    name: "ProvinceID",
                  }}
                >
                  <option aria-label="None" value="">
                    กรุณาเลือกจังหวัด
                  </option>
                  {provinces.map((item: ProvincesInterface) => (
                    <option value={item.ID} key={item.ID}>
                      {item.Name}
                    </option>
                  ))}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={6}>
          <FormControl fullWidth variant="outlined">
              <p>หลักสูตร</p>
              <Select
                native
                value={student.ProgramID + ""}
                onChange={handleChange}
                inputProps={{
                  name: "ProgramID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกหลักสูตรการศึกษา
                </option>
                {program.map((item: ProgramInterface) => (
                <option value={item.ID} key={item.ID}>
                  {item.Program_name}
                </option>
                ))}
              </Select>
            </FormControl>
          </Grid>


          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>บทบาท</p>
              <Select
                native
                value={student.RoleID + ""}
                onChange={handleChange}
                inputProps={{
                  name: "RoleID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกบทบาท
                </option>
                {role.map((item: RoleInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Role_name}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
  
            <Grid item xs={12}>
              <Button
                component={RouterLink}
                to="/students"
                variant="contained"
                color="inherit"
              >
                กลับ
              </Button>
              <Button
                style={{ float: "right" }}
                onClick={submit}
                variant="contained"
                color="primary"
              >
                บันทึก
              </Button>
            </Grid>
          </Grid>
        </Paper>
      </Container>
    );
  }
  
  export default StudentCreate;